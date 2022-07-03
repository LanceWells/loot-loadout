package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.opentelemetry.io/otel"
)

var _ characterimage.CharacterImageDatabase = CharacterImageDB{}

// CharacterImageDB is a psql implementation of a CharacterImage-related database.
type CharacterImageDB struct {
	l  *log.Logger
	db *sql.DB
}

// NewImageDatabase creates a new instance of a CharacterImageDB.
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
	if err != nil {
		l.Println(err)
		return nil, err
	}

	return &CharacterImageDB{
		l:  l,
		db: db,
	}, nil
}

// AddBody adds a new body to the database. This is the base action for all other part additions.
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

// ListBodies returns a list of all of the bodies in the database.
// TODO: Make this accept a filter.
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

// AddDynamicMapping creates a new dynamic mapping for a given part on a given body. We should have
// at least one of these per-part, per-body that has that part. For example, if we were to create a
// new humanoid body type, we should include one of these for each appendage. If we made a reptilian
// creature instead, this list may include a tail in addition to those other appendages.
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
		BodyTypeID: bodyTypeID,
		PartType:   partType,
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
				X:                    int16(x),
				Y:                    int16(y),
			}
		}
	}

	colorStrings, err := db.insertMissingColors(ctx, colors)

	// Actually insert the mapping's pixels at this point.
	for _, c := range colorStrings {
		pixel := pixels[c.Hexstring]
		pixel.ColorStringID = c.ID

		err = pixel.Insert(ctx, db.db, boil.Infer())
		if err != nil {
			db.l.Println("error inserting mapping pixel:", err)
			return "", err
		}
	}

	return strconv.Itoa(mapping.ID), nil
}

// AddStatic creates a new static part. This static part is immutable, but may have its positioning,
// rotation, and color palette changed in a given animation frame.
func (db CharacterImageDB) AddStatic(
	ctx context.Context,
	s *api.Static,
	bodyID characterimage.ID,
) (id characterimage.ID, err error) {
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

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(s.Image)))
	_, err = base64.StdEncoding.Decode(dst, []byte(s.Image))
	if err != nil {
		fmt.Println("decoding error:", err)
		return "", err
	}

	staticType := models.StaticPartType(api.StaticPartType_name[int32(s.Part)])
	if err = staticType.IsValid(); err != nil {
		err = fmt.Errorf("invalid static part type: %v", s.Part)
		db.l.Println(err)
		return "", err
	}

	bodyIDVal, err := strconv.Atoi(bodyID)
	if err != nil {
		db.l.Println("error converting ID value for body:", err)
		return "", err
	}

	static := models.StaticPart{
		BodyTypeID:  bodyIDVal,
		DisplayName: s.DisplayName,
		PartType:    staticType,
	}

	err = static.Insert(ctx, db.db, boil.Infer())
	if err != nil {
		fmt.Println("error inserting static:", err)
		return "", err
	}

	staticImage := models.StaticPartImage{
		X:            int16(s.Anchor.X),
		Y:            int16(s.Anchor.Y),
		ImageBytes:   dst,
		StaticPartID: static.ID,
	}

	err = staticImage.Insert(ctx, db.db, boil.Infer())
	if err != nil {
		fmt.Println("error inserting staticImage:", err)
		return "", err
	}

	return "", nil
}

// ListStatics returns a list of all statics in the DB.
// TODO: Make this accept a filter.
func (db CharacterImageDB) ListStatics(
	ctx context.Context,
	f *characterimage.StaticPartsFilter,
) (map[string]*api.Static, error) {
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

	statics, err := models.StaticParts(
		qm.Load(models.StaticPartRels.StaticPartImage),
	).All(ctx, db.db)
	if err != nil {
		fmt.Println("error getting statics:", err)
		return nil, err
	}

	apiStatics := make(map[string]*api.Static, len(statics))
	for _, v := range statics {
		staticType, ok := api.StaticPartType_value[v.PartType.String()]
		if !ok {
			err = fmt.Errorf("staticType %v does not exist", v.PartType.String())
			fmt.Println("error getting statics:", err)
			return nil, err
		}

		apiStatics[strconv.Itoa(v.ID)] = &api.Static{
			DisplayName: v.DisplayName,
			Part:        api.StaticPartType(staticType),
			Image:       v.R.StaticPartImage.ImageBytes,
			Anchor: &api.Coordinates{
				X: uint32(v.R.StaticPartImage.X),
				Y: uint32(v.R.StaticPartImage.Y),
			},
		}
	}

	return apiStatics, nil
}

// AddDynamic inserts a new dynamic part into the DB. This is effectively a "color map" that matches
// up with a given dynamic part mapping. The position of each pixel correlates to the virtual space
// of that pixel on a mapping.
func (db CharacterImageDB) AddDynamic(
	ctx context.Context,
	d *api.Dynamic,
	bodyID characterimage.ID,
	imageBytes []byte,
) (id characterimage.ID, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "AddDynamic")
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
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(imageBytes)))
	_, err = base64.StdEncoding.Decode(dst, []byte(imageBytes))
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

	// Validate the part type. If it's not a valid enum, we can't insert it.
	partType := models.DynamicPartType(api.DynamicPartType_name[int32(d.Part)])
	err = partType.IsValid()
	if err != nil {
		err = fmt.Errorf("invalid dynamic part type: %v", d.Part)
		db.l.Println(err)
		return "", err
	}

	// Grab the part mapping for this part. We'll use it to validate the incoming png, and we need a
	// ref to it for our dynamic part.
	mapping, err := models.DynamicPartMappings(
		qm.Where(fmt.Sprintf("%v = ?", models.DynamicPartMappingColumns.BodyTypeID), bodyID),
		qm.And(fmt.Sprintf("%v = ?", models.DynamicPartMappingColumns.PartType), partType),
	).One(ctx, db.db)
	if err != nil {
		err = fmt.Errorf("could not find matching part mapping:%v", err)
		db.l.Println(err)
		return "", err
	}

	// Insert the dynamic part.
	dynamic := models.DynamicPart{
		DynamicPartMappingID: mapping.ID,
		DisplayName:          d.DisplayName,
		PartType:             partType,
	}

	err = dynamic.Insert(ctx, db.db, boil.Infer())
	if err != nil {
		err = fmt.Errorf("error inserting dynamic part:%v", err)
		db.l.Println(err)
		return "", err
	}

	// Ensure that any of our incoming png pixels have a matching mapping pixel. We do this to ensure
	// that the provided input doesn't accidentally contain info that we won't use.
	mappingPixels, err := models.DynamicPartMappingPixels(
		qm.Where(fmt.Sprintf("%v = ?", models.DynamicPartMappingPixelColumns.DynamicPartMappingID), mapping.ID),
	).All(ctx, db.db)
	if err != nil {
		err = fmt.Errorf("error getting mapping pixels:%v", err)
		db.l.Println(err)
		return "", err
	}

	mappingPixelMap := make(
		map[image.Point]*models.DynamicPartMappingPixel,
		len(mappingPixels),
	)
	for _, p := range mappingPixels {
		mappingPixelMap[image.Point{
			X: int(p.X),
			Y: int(p.Y),
		}] = p
	}

	// Go over the incoming image. We want to get info about its layout s.t. we have all of its colors
	// normalized, and we have an auto-crop space to crop our thumbnail to.
	var cropRect *image.Rectangle = nil
	colors := []string{}
	pixels := map[string]*models.DynamicPartPixel{}
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

			// If we haven't yet encountered a pixel, set this as our top-left.
			if cropRect == nil {
				cropRect = &image.Rectangle{
					Min: image.Point{
						X: x,
						Y: y,
					},
				}
			}

			// max is always just the last pixel we've encountered that isn't transparent.
			cropRect.Max = image.Point{
				X: x,
				Y: y,
			}

			_, ok := mappingPixelMap[image.Point{
				X: x,
				Y: y,
			}]
			if !ok {
				err = fmt.Errorf(
					"provided dynamic pixel at %v/%v does not have a matching pixel for type %v",
					x,
					y,
					partType,
				)
				db.l.Println(err)
				return "", err
			}

			hexString := fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
			colors = append(colors, hexString)
			pixels[hexString] = &models.DynamicPartPixel{
				DynamicPartID: dynamic.ID,
				X:             int16(x),
				Y:             int16(y),
			}
		}
	}

	colorStrings, err := db.insertMissingColors(ctx, colors)
	if err != nil {
		return "", err
	}

	croppedImage, err := db.cropImage(ctx, img, *cropRect)
	if err != nil {
		return "", err
	}

	cropBuf := bytes.Buffer{}
	err = png.Encode(&cropBuf, croppedImage)
	if err != nil {
		fmt.Println("error cropping thumbnail:", err)
		return "", err
	}

	// Insert the auto-cropped thumbnail.
	thumbnail := models.DynamicPartThumbnail{
		DynamicPartID: dynamic.ID,
		ImageBytes:    cropBuf.Bytes(),
	}

	err = thumbnail.Insert(ctx, db.db, boil.Infer())
	if err != nil {
		err = fmt.Errorf("error inserting dynamic thumbnail:%v", err)
		db.l.Println(err)
		return "", err
	}

	for _, c := range colorStrings {
		pixel := pixels[c.Hexstring]
		pixel.ColorStringID = c.ID

		err = pixel.Insert(ctx, db.db, boil.Infer())
		if err != nil {
			db.l.Println("error inserting dynamic pixel:", err)
			return "", err
		}
	}

	return strconv.Itoa(dynamic.ID), nil
}

// ListDynamics returns a list of all dynamic parts in the DB.
// TODO: Make this accept a filter.
func (db CharacterImageDB) ListDynamics(
	ctx context.Context,
	f *characterimage.DynamicPartsFilter,
) (map[string]*api.Dynamic, map[string][]byte, error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "ListDynamics")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		span.End()
	}()

	dynamics, err := models.DynamicParts(
		qm.Load(models.DynamicPartRels.DynamicPartThumbnail),
	).All(ctx, db.db)
	if err != nil {
		fmt.Println("error getting dynamics:", err)
		return nil, nil, err
	}

	apiDynamics := make(map[string]*api.Dynamic, len(dynamics))
	apiThumbnails := make(map[string][]byte, len(dynamics))
	for _, v := range dynamics {
		dynamicPartType, ok := api.DynamicPartType_value[v.PartType.String()]
		if !ok {
			err = fmt.Errorf("dyanmicType %v does not exist", v.PartType.String())
			fmt.Println("error getting dynamics:", err)
			return nil, nil, err
		}

		id := strconv.Itoa(v.ID)
		apiDynamics[id] = &api.Dynamic{
			DisplayName: v.DisplayName,
			Part:        api.DynamicPartType(dynamicPartType),
		}

		apiThumbnails[id] = v.R.DynamicPartThumbnail.ImageBytes
	}

	return apiDynamics, apiThumbnails, nil
}

// AddAnimation creates a new animation in the DB.
func (db CharacterImageDB) AddAnimation(
	ctx context.Context,
	a *api.Animation,
	bodyID characterimage.ID,
) (characterimage.ID, error) {
	return "", nil
}

// ListAnimations returns a list of all animations in the DB.
// TODO: Make this accept a filter.
func (db CharacterImageDB) ListAnimations(
	ctx context.Context,
	f *characterimage.AnimationsFilter,
) (map[string]*api.Animation, error) {
	return nil, nil
}

// AddFrame inserts a new frame into the given animation in the DB.
func (db CharacterImageDB) AddFrame(
	ctx context.Context,
	f *api.Frame,
	animationID characterimage.ID,
) (characterimage.ID, error) {
	return "", nil
}

// AddProp inserts a new prop into the DB.
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
		PartType:    propType,
	}

	err = prop.Insert(ctx, db.db, boil.Infer())
	if err != nil {
		fmt.Println("error inserting prop:", err)
		return "", err
	}

	propImage := models.PropImage{
		X:          int16(p.Anchor.X),
		Y:          int16(p.Anchor.Y),
		ImageBytes: dst,
		PropID:     prop.ID,
	}

	err = propImage.Insert(ctx, db.db, boil.Infer())
	if err != nil {
		fmt.Println("error inserting propType:", err)
		return "", err
	}

	return strconv.Itoa(prop.ID), nil
}

// ListProps returns a list of all props in the DB.
// TODO: Make this accept a filter.
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

	props, err := models.Props(
		qm.Load(models.PropRels.PropImage),
	).All(ctx, db.db)
	if err != nil {
		fmt.Println("error getting props:", err)
		return nil, err
	}

	apiProps := make(map[string]*api.Prop, len(props))
	for _, v := range props {
		partType, ok := api.PropType_value[v.PartType.String()]
		if !ok {
			err = fmt.Errorf("proptype %v does not exist in API", v.PartType.String())
			fmt.Println("error getting props:", err)
			return nil, err
		}

		apiProps[strconv.Itoa(v.ID)] = &api.Prop{
			DisplayName: v.DisplayName,
			Prop:        api.PropType(partType),
			Image:       v.R.PropImage.ImageBytes,
			Anchor: &api.Coordinates{
				X: uint32(v.R.PropImage.X),
				Y: uint32(v.R.PropImage.Y),
			},
		}
	}

	return apiProps, nil
}

// insertMissingColors is a helper method that accepts a list of colors, then returns the DB color
// strings associated with those colors. This method will also insert any colors into the DB that do
// not already exist.
func (db CharacterImageDB) insertMissingColors(
	ctx context.Context,
	colors []string,
) (models.ColorStringSlice, error) {
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
		return nil, err
	}

	// This feels a little extra, but it could save us a lot of loops, so why not?
	if len(colors) != len(colorStrings) {
		colorStringMap := map[string]bool{}
		for _, c := range colorStrings {
			colorStringMap[c.Hexstring] = true
		}

		for _, c := range colors {
			if _, ok := colorStringMap[c]; !ok {
				colorString := &models.ColorString{
					Hexstring: c,
				}

				err = colorString.Insert(ctx, db.db, boil.Infer())
				if err != nil {
					db.l.Println("error inserting color", err)
					return nil, err
				}

				colorStrings = append(colorStrings, colorString)
			}
		}
	}

	return colorStrings, nil
}

func (db CharacterImageDB) cropImage(
	ctx context.Context,
	img image.Image,
	rect image.Rectangle,
) (image.Image, error) {
	type canSubImage interface {
		SubImage(r image.Rectangle) image.Image
	}

	croppable, ok := img.(canSubImage)
	if !ok {
		return img, nil
	}

	return croppable.SubImage(rect), nil
}
