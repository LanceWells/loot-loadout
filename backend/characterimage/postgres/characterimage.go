package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"

	api "github.com/lantspants/lootloadout/api/characterimage/v1"
	"github.com/lantspants/lootloadout/backend/characterimage"
	"github.com/lantspants/lootloadout/backend/characterimage/postgres/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.opentelemetry.io/otel"
)

/*
dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
*/

type CharacterImageDB struct {
	l  *log.Logger
	db *sql.DB
}

func NewImageDatabase(
	ctx context.Context,
	l *log.Logger,
) (*CharacterImageDB, error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "NewImageDatabase")
	defer span.End()

	dbName, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		err := "POSTGRES_DB env var is not set"
		l.Println(err)
		return nil, errors.New(err)
	}

	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		err := "POSTGRES_USER env var is not set"
		l.Println(err)
		return nil, errors.New(err)
	}

	pass, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		err := "POSTGRES_PASSWORD env var is not set"
		l.Println(err)
		return nil, errors.New(err)
	}

	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		err := "POSTGRES_PORT env var is not set"
		l.Println(err)
		return nil, errors.New(err)
	}

	dsn := strings.Join([]string{
		"host=character-image-db",
		"sslmode=disable",
		fmt.Sprintf("user=%s", user),
		fmt.Sprintf("password=%s", pass),
		fmt.Sprintf("dbname=%s", dbName),
		fmt.Sprintf("port=%s", port),
	}, " ")

	db, err := sql.Open("postgres", dsn)
	// d := postgres.Open(dsn)
	// db, err := gorm.Open(d, &gorm.Config{})
	if err != nil {
		l.Println(err)
		return nil, err
	}

	return &CharacterImageDB{
		l:  l,
		db: db,
	}, nil
}

func (db CharacterImageDB) AddBody(
	ctx context.Context,
	b *api.Body,
) (characterimage.ID, error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "AddBody")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		span.End()
	}()

	body := models.BodyType{
		DisplayName: b.DisplayName,
	}

	err = body.Insert(ctx, db.db, boil.Infer())
	if err != nil {
		fmt.Println("error inserting body:", err)
		return "", err
	}

	return fmt.Sprintf("%v", body.ID), nil
}

func (db CharacterImageDB) ListBodies(
	ctx context.Context,
	f *characterimage.BodyTypesFilter,
) (map[string]*api.Body, error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "ListBodies")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		span.End()
	}()

	bodies, err := models.BodyTypes().All(ctx, db.db)
	if err != nil {
		fmt.Println("error getting bodies:", err)
		return nil, err
	}

	apiBodies := make(map[string]*api.Body, len(bodies))
	for _, v := range bodies {
		apiBodies[strconv.Itoa(v.ID)] = &api.Body{
			DisplayName: v.DisplayName,
		}
	}

	return apiBodies, nil
}

func (db CharacterImageDB) AddDynamicMapping(
	ctx context.Context,
	m *api.DynamicMapping,
	bodyID characterimage.ID,
) (id characterimage.ID, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "AddDynamicMapping")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
		span.End()
	}()

	// First, ensure that the provided PNG is properly base64-encoded, and that it is also a valid
	// .png image.
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(m.Image)))
	_, err = base64.StdEncoding.Decode(dst, []byte(m.Image))
	if err != nil {
		fmt.Println("decoding error:", err)
		return "", err
	}

	r := bytes.NewReader(dst)
	img, err := png.Decode(r)
	if err != nil {
		fmt.Println("invalid png:", err)
		return "", err
	}

	// Validate the enums.
	bodyTypeID, err := strconv.Atoi(bodyID)
	if err != nil {
		err := fmt.Errorf("provided ID %v is invalid, %v", bodyID, err)
		fmt.Println(err)
		return "", err
	}

	partType := models.DynamicPartType(api.DynamicPartType_name[int32(m.Part)])
	if err = partType.IsValid(); err != nil {
		err = fmt.Errorf("invalid dynamic part type: %v", m.Part)
		db.l.Println(err)
		return "", err
	}

	// Insert the base mapping. This needs to happen before we insert any of the pixels. This should
	// be rolled-back if we have any errors in this transaction.
	mapping := models.DynamicPartMapping{
		BodyTypeID: null.IntFrom(bodyTypeID),
		PartType:   models.NullDynamicPartTypeFrom(partType),
	}

	err = mapping.Insert(ctx, db.db, boil.Infer())
	if err != nil {
		fmt.Println("error inserting dynamic mapping", err)
		return "", err
	}

	// Make a list of pixel objects first. This lets us evaluate the .png image, which lets us get a
	// list of all of the colors that we need a reference to. From there, we can get the IDs for all
	// of those pixels (inserting if any given pixel does not yet exist).
	colors := []string{}
	pixels := map[string]*models.DynamicPartMappingPixel{}
	b := img.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			c := color.NRGBAModel.Convert(b.RGBA64At(x, y)).(color.NRGBA)
			r := c.R
			g := c.G
			b := c.B
			a := c.A

			// Do not save transparent pixels.
			if a == 0 {
				continue
			}

			hexString := fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
			colors = append(colors, hexString)
			pixels[hexString] = &models.DynamicPartMappingPixel{
				DynamicPartMappingID: mapping.ID,
				X:                    null.Int16From(int16(x)),
				Y:                    null.Int16From(int16(y)),
			}
		}
	}

	colorInterfaces := make([]interface{}, len(colors))
	for _, c := range colors {
		colorInterfaces = append(colorInterfaces, c)
	}

	colorStrings, err := models.ColorStrings(
		qm.WhereIn(
			fmt.Sprintf("%v in ?", models.ColorStringColumns.Hexstring),
			colorInterfaces...,
		),
	).All(ctx, db.db)
	if err != nil {
		db.l.Println("error getting colorstrings:", err)
		return "", err
	}

	// This feels a little extra, but it could save us a lot of loops, so why not?
	if len(colors) != len(colorStrings) {
		colorStringMap := map[string]bool{}
		for _, c := range colorStrings {
			colorStringMap[c.Hexstring.String] = true
		}

		for _, c := range colors {
			if _, ok := colorStringMap[c]; !ok {
				colorString := &models.ColorString{
					Hexstring: null.StringFrom(c),
				}

				err = colorString.Insert(ctx, db.db, boil.Infer())
				if err != nil {
					db.l.Println("error inserting color", err)
					return "", err
				}

				colorStrings = append(colorStrings, colorString)
			}
		}
	}

	// Actually insert the mapping's pixels at this point.
	for _, c := range colorStrings {
		pixel := pixels[c.Hexstring.String]
		pixel.ColorStringID = c.ID

		err = pixel.Insert(ctx, db.db, boil.Infer())
		if err != nil {
			db.l.Println("error inserting mapping pixel:", err)
			return "", err
		}
	}

	return strconv.Itoa(mapping.ID), nil
}

func (db CharacterImageDB) AddStatic(
	ctx context.Context,
	s *api.Static,
) (characterimage.ID, error) {
	return "", nil
}

func (db CharacterImageDB) ListStatics(
	ctx context.Context,
	f *characterimage.StaticPartsFilter,
) (map[string]*api.Static, error) {
	return nil, nil
}

func (db CharacterImageDB) AddDynamic(
	ctx context.Context,
	d *api.Dynamic,
	bodyID characterimage.ID,
) (characterimage.ID, error) {
	return "", nil
}

func (db CharacterImageDB) ListDynamics(
	ctx context.Context,
	f *characterimage.DynamicPartsFilter,
) (map[string]*api.Dynamic, error) {
	return nil, nil
}

func (db CharacterImageDB) AddAnimation(
	ctx context.Context,
	a *api.Animation,
	bodyID characterimage.ID,
) (characterimage.ID, error) {
	return "", nil
}

func (db CharacterImageDB) ListAnimations(
	ctx context.Context,
	f *characterimage.AnimationsFilter,
) (map[string]*api.Animation, error) {
	return nil, nil
}

func (db CharacterImageDB) AddFrame(
	ctx context.Context,
	f *api.Frame,
	animationID characterimage.ID,
) (characterimage.ID, error) {
	return "", nil
}

func (db CharacterImageDB) AddProp(
	ctx context.Context,
	p *api.Prop,
) (characterimage.ID, error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "CreateProp")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		span.End()
	}()

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(p.Image)))
	_, err = base64.StdEncoding.Decode(dst, []byte(p.Image))
	if err != nil {
		fmt.Println("decoding error:", err)
		return "", err
	}

	propType := models.PropType(api.PropType_name[int32(p.Prop)])
	if err = propType.IsValid(); err != nil {
		err = fmt.Errorf("invalid prop type: %v", p.Prop)
		db.l.Println(err)
		return "", err
	}

	prop := models.Prop{
		DisplayName: p.DisplayName,
		PartType:    models.NullPropTypeFrom(propType),
	}

	err = prop.Insert(ctx, db.db, boil.Infer())
	if err != nil {
		fmt.Println("error inserting prop:", err)
		return "", err
	}

	propImage := models.PropImage{
		X:          null.Int16From(int16(p.Anchor.X)),
		Y:          null.Int16From(int16(p.Anchor.Y)),
		ImageBytes: null.BytesFrom(dst),
		PropID:     prop.ID,
	}

	err = propImage.Insert(ctx, db.db, boil.Infer())
	if err != nil {
		fmt.Println("error inserting propType:", err)
		return "", err
	}

	return strconv.Itoa(prop.ID), nil
}

func (db CharacterImageDB) ListProps(
	ctx context.Context,
	f *characterimage.PropsFilter,
) (map[string]*api.Prop, error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "ListProps")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		span.End()
	}()

	props, err := models.Props(qm.Load(models.PropRels.PropImage)).All(ctx, db.db)
	if err != nil {
		fmt.Println("error getting props:", err)
		return nil, err
	}

	// apiProps := make([]*api.Prop, len(props))
	apiProps := make(map[string]*api.Prop, len(props))
	for _, v := range props {
		if !v.PartType.Valid {
			err := errors.New("PropType on returned prop is invalid")
			fmt.Println("error getting props:", err)
			return nil, err
		}

		partType, ok := api.PropType_value[v.PartType.Val.String()]
		if !ok {
			err = fmt.Errorf("proptype %v does not exist in API", partType)
			fmt.Println("error getting props:", err)
			return nil, err
		}

		apiProps[strconv.Itoa(v.ID)] = &api.Prop{
			DisplayName: v.DisplayName,
			Prop:        api.PropType(partType),
			Image:       v.R.PropImage.ImageBytes.Bytes,
			Anchor: &api.Coordinates{
				X: uint32(v.R.PropImage.X.Int16),
				Y: uint32(v.R.PropImage.Y.Int16),
			},
		}
	}

	return apiProps, nil
}

// func (db CharacterImageDB) CreateBodyType(
// 	ctx context.Context,
// 	bt *pb.BodyType,
// ) (characterimage.ID, error) {
// 	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "CreateBodyType")
// 	defer span.End()

// 	// Create the body type first. Everything needs to be able to reference this eventually.
// 	body := &model.BodyType{
// 		DisplayName: bt.DisplayName,
// 	}

// 	tx := db.db.Create(body)
// 	if tx.Error != nil {
// 		db.l.Printf("error creating new body type: %v", tx.Error)
// 		return "", tx.Error
// 	}

// 	// Go through each of the dynamic part mappings and ensure that we are able to reference them.
// 	dpms := []model.DynamicPartMapping{}
// 	for dpm := range bt.PartMaps {
// 		dpms = append(dpms, model.DynamicPartMapping{
// 			BodyTypeID: sql.NullInt64{
// 				Int64: int64(body.ID),
// 				Valid: true,
// 			},
// 			PartType: sql.NullString{
// 				String: dpm,
// 				Valid:  true,
// 			},
// 		})
// 	}

// 	tx = db.db.Create(dpms)
// 	if tx.Error != nil {
// 		db.l.Printf("error creating dynamic part mappings: %v", tx.Error)
// 		return "", tx.Error
// 	}

// 	// Create a temporary map to look up the part type in order to get its newly-created ID.
// 	dpmLookup := map[string]int32{}
// 	for _, dpm := range dpms {
// 		dpmLookup[dpm.PartType.String] = dpm.ID
// 	}

// 	// Save the DPM pixels for each part type. Note that this is being saved in a way that we're
// 	// keeping the hex value for each pixel as a raw string. We do this because the DB is structured
// 	// s.t. our color hexes are normalized. This is intentional as the DB should only host ~64 colors
// 	// in total (not counting alpha-variants), so storing N strings where N = the number of pixels
// 	// across all images is dramatically more expensive than saving that number of smallints.
// 	//
// 	// The solution here is that we still evaluate the bytes for each image, but instead of creating
// 	// each pixel we instead stash those pixels away and get the raw hex value for each color. We can
// 	// then fetch all of the IDs associated with those colors before attempting to save these pixels.
// 	dpmsPixels := []struct {
// 		model.DynamicPartMappingPixel
// 		HexValue string
// 	}{}

// 	for k, v := range bt.PartMaps {
// 		r := bytes.NewReader(v)
// 		img, err := png.Decode(r)
// 		if err != nil {
// 			return "", err
// 		}

// 		b := img.Bounds()
// 		for x := b.Min.X; x < b.Max.X; x++ {
// 			for y := b.Min.Y; y < b.Min.Y; y++ {
// 				r, g, b, a := b.At(x, y).RGBA()

// 				// Do not save transparent pixels.
// 				if a >= 255 {
// 					continue
// 				}

// 				hexString := fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
// 				dpmID, ok := dpmLookup[k]
// 				if !ok {
// 					err := "failure when finding matching part type in dpm lookup"
// 					db.l.Printf(err)
// 					return "", errors.New(err)
// 				}

// 				dpmsPixels = append(dpmsPixels, struct {
// 					model.DynamicPartMappingPixel
// 					HexValue string
// 				}{
// 					DynamicPartMappingPixel: model.DynamicPartMappingPixel{
// 						DynamicPartMappingID: dpmID,
// 					},
// 					HexValue: hexString,
// 				})
// 			}
// 		}
// 	}

// 	// Get the color values that we need indices for.
// 	hexStrings := []model.ColorString{}
// 	for _, p := range dpmsPixels {
// 		hexStrings = append(hexStrings, model.ColorString{
// 			Hexstring: sql.NullString{
// 				String: p.HexValue,
// 				Valid:  true,
// 			},
// 		})
// 	}

// 	tx = db.db.Clauses(
// 		clause.OnConflict{DoNothing: true},
// 	).Create(hexStrings)
// 	if tx.Error != nil {
// 		db.l.Printf(tx.Error.Error())
// 		return "", tx.Error
// 	}

// 	colorStringIDs := map[string]int{}
// 	for _, v := range hexStrings {
// 		colorStringIDs[v.Hexstring.String] = int(v.ID)
// 	}

// 	// Now add those indices back into the pixels we've created temp structs for.
// 	pixels := []model.DynamicPartMappingPixel{}
// 	for _, p := range dpmsPixels {
// 		colorStringID, ok := colorStringIDs[p.HexValue]
// 		if !ok {
// 			err := fmt.Sprintf("failure getting color string %v", p.HexValue)
// 			db.l.Printf(err)
// 			return "", errors.New(err)
// 		}

// 		pixels = append(pixels, model.DynamicPartMappingPixel{
// 			ColorStringID:        int32(colorStringID),
// 			DynamicPartMappingID: p.DynamicPartMappingID,
// 			X:                    p.X,
// 			Y:                    p.Y,
// 		})
// 	}

// 	tx = db.db.Create(pixels)
// 	if tx.Error != nil {
// 		db.l.Printf(tx.Error.Error())
// 		return "", tx.Error
// 	}

// 	id := fmt.Sprintf("%v", body.ID)
// 	return id, nil
// }

// func (r *CharacterImageDB) GetBodyType(
// 	ctx context.Context,
// 	id characterimage.ID,
// ) (*pb.BodyType, error) {
// 	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "CreateBodyType")
// 	defer span.End()

// 	idVal, err := strconv.Atoi(id)
// 	if err != nil {
// 		r.l.Printf("error converting ID to int: %v", err)
// 		return nil, err
// 	}

// 	body := &model.BodyType{
// 		ID: int32(idVal),
// 	}

// 	tx := r.db.First(body)
// 	if tx.Error != nil {
// 		r.l.Printf("error fetching bodyType from database: %v", tx.Error)
// 		return nil, tx.Error
// 	}

// 	return nil, nil
// }

// func (r *CharacterImageDB) ListBodies(
// 	context.Context,
// 	*characterimage.BodyTypesFilter,
// ) ([]pb.BodyType, error) {
// 	return nil, nil
// }

// func (r *CharacterImageDB) CreateStaticPart(
// 	context.Context,
// 	*pb.StaticPart,
// ) (characterimage.ID, error) {
// 	return "", nil
// }

// func (r *CharacterImageDB) GetStaticPart(
// 	context.Context,
// 	characterimage.ID,
// ) (*pb.StaticPart, error) {
// 	return nil, nil
// }

// func (r *CharacterImageDB) ListStaticParts(
// 	context.Context,
// 	*characterimage.StaticPartsFilter,
// ) ([]pb.StaticPart, error) {
// 	return nil, nil
// }

// func (r *CharacterImageDB) CreateDynamicPart(
// 	context.Context,
// 	*pb.DynamicPart,
// ) (characterimage.ID, error) {
// 	return "", nil
// }

// func (r *CharacterImageDB) GetDynamicPart(
// 	context.Context,
// 	characterimage.ID,
// ) (*pb.DynamicPart, error) {
// 	return nil, nil
// }

// func (r *CharacterImageDB) ListDynamicParts(
// 	context.Context,
// 	*characterimage.DynamicPartsFilter,
// ) ([]pb.DynamicPart, error) {
// 	return nil, nil
// }

// func (r *CharacterImageDB) CreateProp(
// 	context.Context,
// 	*pb.Prop,
// ) (characterimage.ID, error) {
// 	return "", nil
// }

// func (r *CharacterImageDB) GetProp(
// 	context.Context,
// 	characterimage.ID,
// ) (*pb.Prop, error) {
// 	return nil, nil
// }

// func (r *CharacterImageDB) ListProps(
// 	context.Context,
// 	*characterimage.PropsFilter,
// ) ([]pb.Prop, error) {
// 	return nil, nil
// }

// func (r *CharacterImageDB) CreateAnimation(
// 	context.Context,
// 	*pb.Animation,
// ) (characterimage.ID, error) {
// 	return "", nil
// }

// func (r *CharacterImageDB) GetAnimation(
// 	context.Context,
// 	characterimage.ID,
// ) (*pb.Animation, error) {
// 	return nil, nil
// }

// func (r *CharacterImageDB) ListAnimations(
// 	context.Context,
// 	*characterimage.AnimationsFilter,
// ) ([]pb.BodyType, error) {
// 	return nil, nil
// }
