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
) (id characterimage.ID, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "AddBody")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

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

	err = body.Insert(ctx, tx, boil.Infer())
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
) (apiBodies map[string]*api.Body, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "ListBodies")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		span.End()
	}()

	bodies, err := models.BodyTypes().All(ctx, tx)
	if err != nil {
		fmt.Println("error getting bodies:", err)
		return nil, err
	}

	apiBodies = make(map[string]*api.Body, len(bodies))
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
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

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
	bodyIDVal, err := strconv.Atoi(bodyID)
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
		BodyTypeID: bodyIDVal,
		PartType:   partType,
	}

	err = mapping.Insert(ctx, tx, boil.Infer())
	if err != nil {
		fmt.Println("error inserting dynamic mapping", err)
		return "", err
	}

	// Make a list of pixel objects first. This lets us evaluate the .png image, which lets us get a
	// list of all of the colors that we need a reference to. From there, we can get the IDs for all
	// of those pixels (inserting if any given pixel does not yet exist).
	colors := []string{}
	pixels := map[string]*models.DynamicPartMappingPixel{}
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
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

	colorStrings, err := insertMissingColors(ctx, colors, tx)

	// Actually insert the mapping's pixels at this point.
	for _, c := range colorStrings {
		pixel := pixels[c.Hexstring]
		pixel.ColorStringID = c.ID

		err = pixel.Insert(ctx, tx, boil.Infer())
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
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

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

	err = static.Insert(ctx, tx, boil.Infer())
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

	err = staticImage.Insert(ctx, tx, boil.Infer())
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
) (apiStatics map[string]*api.Static, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "ListProps")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		span.End()
	}()

	statics, err := models.StaticParts(
		qm.Load(models.StaticPartRels.StaticPartImage),
	).All(ctx, tx)
	if err != nil {
		fmt.Println("error getting statics:", err)
		return nil, err
	}

	apiStatics = make(map[string]*api.Static, len(statics))
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
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

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
	).One(ctx, tx)
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

	err = dynamic.Insert(ctx, tx, boil.Infer())
	if err != nil {
		err = fmt.Errorf("error inserting dynamic part:%v", err)
		db.l.Println(err)
		return "", err
	}

	// Ensure that any of our incoming png pixels have a matching mapping pixel. We do this to ensure
	// that the provided input doesn't accidentally contain info that we won't use.
	mappingPixels, err := models.DynamicPartMappingPixels(
		qm.Where(fmt.Sprintf("%v = ?", models.DynamicPartMappingPixelColumns.DynamicPartMappingID), mapping.ID),
	).All(ctx, tx)
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

	colors := []string{}
	pixels := map[string][]*models.DynamicPartPixel{}
	bounds := img.Bounds()

	// Go over the incoming image. We want to get info about its layout s.t. we have all of its colors
	// normalized, and we have an auto-crop space to crop our thumbnail to. Note that this strategy
	// only works if our image isn't completely transparent. That's a fix TODO: in future.
	var cropRect *image.Rectangle = &image.Rectangle{
		Min: image.Pt(bounds.Max.X, bounds.Max.Y),
		Max: image.Pt(bounds.Min.X, bounds.Min.Y),
	}

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			r := c.R
			g := c.G
			b := c.B
			a := c.A

			// Do not save transparent pixels.
			if a == 0 {
				continue
			}

			// Max is always just the last pixel we've encountered that isn't transparent. We add 1 to
			// each because the max points are exclusive, whereas the Min points are inclusive. This
			// ensures that we're referencing the correct row/column.
			if cropRect.Min.X > x {
				cropRect.Min.X = x
			}
			if cropRect.Min.Y > y {
				cropRect.Min.Y = y
			}
			if cropRect.Max.X < (x + 1) {
				cropRect.Max.X = (x + 1)
			}
			if cropRect.Max.Y < (y + 1) {
				cropRect.Max.Y = (y + 1)
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
			pixelListAtColor, ok := pixels[hexString]
			if !ok {
				pixelListAtColor = []*models.DynamicPartPixel{}
			}

			pixels[hexString] = append(pixelListAtColor, &models.DynamicPartPixel{
				DynamicPartID: dynamic.ID,
				X:             int16(x),
				Y:             int16(y),
			})
		}
	}

	colorStrings, err := insertMissingColors(ctx, colors, tx)
	if err != nil {
		fmt.Println("error inserting colors:", err)
		return "", err
	}

	colorStringMap := map[string]int{}
	for _, c := range colorStrings {
		colorStringMap[c.Hexstring] = c.ID
	}

	croppedImage, err := db.cropImage(ctx, img, *cropRect)
	if err != nil {
		fmt.Println("error cropping image:", err)
		return "", err
	}

	cropBuf := bytes.Buffer{}
	err = png.Encode(&cropBuf, croppedImage)
	if err != nil {
		fmt.Println("error cropping thumbnail:", err)
		return "", err
	}

	// cropBytes := cropBuf.Bytes()
	// thumbEncodedLen := base64.StdEncoding.EncodedLen(len(cropBytes))
	// thumbDst := make([]byte, thumbEncodedLen)
	// base64.StdEncoding.Encode(thumbDst, cropBytes)

	// Insert the auto-cropped thumbnail.
	thumbnail := models.DynamicPartThumbnail{
		DynamicPartID: dynamic.ID,
		ImageBytes:    cropBuf.Bytes(),
	}

	err = thumbnail.Insert(ctx, tx, boil.Infer())
	if err != nil {
		err = fmt.Errorf("error inserting dynamic thumbnail:%v", err)
		db.l.Println(err)
		return "", err
	}

	for k, v := range pixels {
		colorStringID, ok := colorStringMap[k]
		if !ok {
			err = fmt.Errorf("could not find color for hex %v", k)
			db.l.Println(err)
			return "", err
		}

		for _, p := range v {
			p.ColorStringID = colorStringID
			err = p.Insert(ctx, tx, boil.Infer())
			if err != nil {
				db.l.Println("error inserting dynamic pixel", err)
				return "", err
			}
		}
	}

	return strconv.Itoa(dynamic.ID), nil
}

// ListDynamics returns a list of all dynamic parts in the DB.
// TODO: Make this accept a filter.
func (db CharacterImageDB) ListDynamics(
	ctx context.Context,
	f *characterimage.DynamicPartsFilter,
) (apiDynamics map[string]*api.Dynamic, apiThumbnails map[string][]byte, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "ListDynamics")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		span.End()
	}()

	dynamics, err := models.DynamicParts(
		qm.Load(models.DynamicPartRels.DynamicPartThumbnail),
	).All(ctx, tx)
	if err != nil {
		fmt.Println("error getting dynamics:", err)
		return nil, nil, err
	}

	apiDynamics = make(map[string]*api.Dynamic, len(dynamics))
	apiThumbnails = make(map[string][]byte, len(dynamics))
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

		// Thumbnails are *technically* not needed, though they should be added with each entry.
		if v.R.DynamicPartThumbnail != nil {
			apiThumbnails[id] = v.R.DynamicPartThumbnail.ImageBytes
		}
	}

	return apiDynamics, apiThumbnails, nil
}

// AddAnimation creates a new animation in the DB.
func (db CharacterImageDB) AddAnimation(
	ctx context.Context,
	a *api.Animation,
	bodyID characterimage.ID,
) (id characterimage.ID, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "AddAnimation")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
		span.End()
	}()

	bodyIDVal, err := strconv.Atoi(bodyID)
	if err != nil {
		db.l.Println("error converting ID value for body:", err)
		return "", err
	}

	// This is technically an array of models.PropType, but sqlboiler converts these to strings.
	partTypes := []string{}
	for _, v := range a.Allowed {
		partType := models.PropType(api.PropType_name[int32(v)])
		err = partType.IsValid()
		if err != nil {
			err = fmt.Errorf("invalid prop type: %v", v)
			db.l.Println(err)
			return "", err
		}

		partTypes = append(partTypes, string(partType))
	}

	animation := models.Animation{
		BodyTypeID:  bodyIDVal,
		DisplayName: a.DisplayName,
		PartType:    partTypes,
	}

	err = animation.Insert(ctx, tx, boil.Infer())
	if err != nil {
		fmt.Println("error inserting animation:", err)
		return "", err
	}

	return strconv.Itoa(animation.ID), nil
}

// ListAnimations returns a list of all animations in the DB.
// TODO: Make this accept a filter.
func (db CharacterImageDB) ListAnimations(
	ctx context.Context,
	f *characterimage.AnimationsFilter,
) (apiAnimations map[string]*api.Animation, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "ListAnimations")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		span.End()
	}()

	animations, err := models.Animations().All(ctx, tx)
	if err != nil {
		fmt.Println("error listing animations:", err)
		return nil, err
	}

	apiAnimations = make(map[string]*api.Animation, len(animations))
	for _, a := range animations {
		allowed := []api.PropType{}
		for _, pt := range a.PartType {
			ptVal, ok := api.PropType_value[pt]
			if !ok {
				err = fmt.Errorf("propType %v does not exist", pt)
				fmt.Println("error getting statics:", err)
				return nil, err
			}

			allowed = append(allowed, api.PropType(ptVal))
		}

		apiAnimations[strconv.Itoa(a.ID)] = &api.Animation{
			DisplayName: a.DisplayName,
			Allowed:     allowed,
		}
	}

	return apiAnimations, nil
}

// AddFrame inserts a new frame into the given animation in the DB.
func (db CharacterImageDB) AddFrame(
	ctx context.Context,
	f *api.Frame,
	animationID characterimage.ID,
	imageBytes []byte,
) (id characterimage.ID, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "AddFrame")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
		span.End()
	}()

	animationIDVal, err := strconv.Atoi(animationID)
	if err != nil {
		err := fmt.Errorf("provided ID %v is invalid, %v", animationID, err)
		fmt.Println(err)
		return "", err
	}

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

	lastFrame, err := models.AnimationFrames(
		qm.Where(fmt.Sprintf("%v = ?", models.AnimationFrameColumns.AnimationID), animationID),
		qm.OrderBy(fmt.Sprintf("%v DESC", models.AnimationFrameColumns.FrameIndex)),
	).One(ctx, tx)
	if err != nil && err != sql.ErrNoRows {
		db.l.Println("error getting last animation frame:", err)
		return "", err
	}

	lastFrameIndex := 0
	if lastFrame != nil {
		lastFrameIndex = lastFrame.FrameIndex + 1
	}

	expression := models.ExpressionType(f.Expression.String())
	err = expression.IsValid()
	if err != nil {
		err = fmt.Errorf("invalid expression: %v", f.Expression.String())
		db.l.Println(err)
		return "", err
	}

	frame := models.AnimationFrame{
		AnimationID: animationIDVal,
		FrameIndex:  lastFrameIndex,
		Expression:  expression,
	}

	err = frame.Insert(ctx, tx, boil.Infer())
	if err != nil {
		db.l.Println("error inserting animation frame:", err)
		return "", err
	}

	propPos := models.AnimationFramePropPosition{
		AnimationFrameID: frame.ID,
		X:                int16(f.PropPositioning.Coordinates.X),
		Y:                int16(f.PropPositioning.Coordinates.Y),
		Rotation:         int16(f.PropPositioning.Rotation),
	}

	propPos.Insert(ctx, tx, boil.Infer())
	if err != nil {
		db.l.Println("error inserting prop pos:", err)
		return "", err
	}

	for k, v := range f.StaticPositioning {
		staticPart := models.StaticPartType(k)
		err = staticPart.IsValid()
		if err != nil {
			err = fmt.Errorf("invalid static part: %v", k)
			db.l.Println(err)
			return "", err
		}

		staticPos := models.AnimationFrameStaticPosition{
			AnimationFrameID: animationIDVal,
			PartType:         staticPart,
			X:                int16(v.Coordinates.X),
			Y:                int16(v.Coordinates.Y),
			Rotation:         int16(v.Rotation),
		}

		err = staticPos.Insert(ctx, tx, boil.Infer())
		if err != nil {
			db.l.Println("error inserting static part pos:", err)
			return "", err
		}
	}

	colors := []string{}
	pixels := map[string]*models.AnimationFramePixel{}
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
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
			pixels[hexString] = &models.AnimationFramePixel{
				AnimationFrameID: animationIDVal,
				X:                int16(x),
				Y:                int16(y),
			}
		}
	}

	colorStrings, err := insertMissingColors(ctx, colors, tx)
	if err != nil {
		return "", err
	}

	for _, c := range colorStrings {
		pixel := pixels[c.Hexstring]
		pixel.ColorStringID = c.ID

		err = pixel.Insert(ctx, tx, boil.Infer())
		if err != nil {
			db.l.Println("error inserting frame pixel:", err)
			return "", err
		}
	}

	return strconv.Itoa(frame.ID), nil
}

// AddProp inserts a new prop into the DB.
func (db CharacterImageDB) AddProp(
	ctx context.Context,
	p *api.Prop,
) (id characterimage.ID, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "CreateProp")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

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

	err = prop.Insert(ctx, tx, boil.Infer())
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

	err = propImage.Insert(ctx, tx, boil.Infer())
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
) (apiProps map[string]*api.Prop, err error) {
	_, span := otel.Tracer("CharacterImageDB").Start(ctx, "ListProps")
	tx, err := db.db.BeginTx(ctx, nil)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("CRITICAL: recovered from panic: %v", r)
			fmt.Println(err)
		}

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		span.End()
	}()

	props, err := models.Props(
		qm.Load(models.PropRels.PropImage),
	).All(ctx, tx)
	if err != nil {
		fmt.Println("error getting props:", err)
		return nil, err
	}

	apiProps = make(map[string]*api.Prop, len(props))
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

func (db CharacterImageDB) GenerateAnimation(
	ctx context.Context,
	animationID string,
	dynamicPartIDs map[*api.DynamicPartType]string,
	staticPartIDs map[*api.StaticPartType]string,
	propID string,
) (animations *characterimage.GeneratedAnimation, err error) {
	return nil, nil
}

// insertMissingColors is a helper method that accepts a list of colors, then returns the DB color
// strings associated with those colors. This method will also insert any colors into the DB that do
// not already exist.
func insertMissingColors(
	ctx context.Context,
	colors []string,
	tx *sql.Tx,
) (models.ColorStringSlice, error) {
	uniqueColors := map[string]bool{}
	for _, c := range colors {
		uniqueColors[c] = true
	}

	colorInterfaces := []interface{}{}
	for k := range uniqueColors {
		colorInterfaces = append(colorInterfaces, k)
	}

	colorStrings, err := models.ColorStrings(
		qm.WhereIn(
			fmt.Sprintf("%v IN ?", models.ColorStringColumns.Hexstring),
			colorInterfaces...,
		),
	).All(ctx, tx)
	if err != nil {
		return nil, err
	}

	if len(uniqueColors) != len(colorStrings) {
		for _, c := range colorStrings {
			uniqueColors[c.Hexstring] = false
		}

		for k, v := range uniqueColors {
			if !v {
				continue
			}

			colorString := &models.ColorString{
				Hexstring: k,
			}

			err = colorString.Insert(ctx, tx, boil.Infer())
			if err != nil {
				return nil, err
			}

			colorStrings = append(colorStrings, colorString)
		}
	}

	return colorStrings, nil
}

// cropImage crops the given image only if that image format type supports cropping.
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
