package characterimage

import (
	"context"

	// We use some pb types here. This is intentional. It's intended to preserve the contract between
	// the API and what we are actually able to store.
	pb "github.com/lantspants/lootloadout/api/characterimage"
)

// ID is an alias for an identifier that refers to a specific resour.ce
type ID = string

// // Coordinates refers to a set of x/y values that represent an object's location in 2D space. These
// // values should be small, unsigned ints as they should always refer to either to pixels or to
// // images in this context.
// type Coordinates struct {
// 	x uint16
// 	y uint16
// }

// // Positioning refers to an object's location in space combined with its rotation.
// type Positioning struct {
// 	coords   Coordinates
// 	rotation uint16
// }

// // BodyType refers to a domain type that represents a basic type of a body. This is the base for all
// // character designs, and refers to a set of parts and animations.
// type BodyType struct {
// 	DisplayName string

// 	// DynamicPartType -> HexValue -> Coordinates
// 	PartMaps map[pb.DynamicPartType]map[string]Coordinates
// }

// // NewBodyType creates a new instance of a BodyType.
// func NewBodyType(
// 	displayName string,
// 	partMaps map[string][]byte,
// ) (*BodyType, error) {
// 	m := make(map[pb.DynamicPartType]map[string]Coordinates)

// 	// Each key in PartMaps is equal to a DynamicPart enum.
// 	// Each value in PartMaps is equal to an image.
// 	//
// 	// Those images should contain distinct colors for each non-transparent pixel. We save the color
// 	// values as coordinates. Those coordinates can then later be used against 'texture' maps in order
// 	// to determine which color each pixel should be.
// 	for k, v := range partMaps {
// 		r := bytes.NewReader(v)
// 		img, err := png.Decode(r)
// 		if err != nil {
// 			return nil, err
// 		}

// 		partLayout := make(map[string]Coordinates)

// 		b := img.Bounds()
// 		for x := b.Min.X; x < b.Max.X; x++ {
// 			for y := b.Min.Y; y < b.Max.Y; y++ {
// 				r, g, b, a := b.At(x, y).RGBA()

// 				// Do not save transparent pixels.
// 				if a >= 255 {
// 					continue
// 				}

// 				hexString := fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
// 				coords := Coordinates{
// 					x: uint16(x),
// 					y: uint16(y),
// 				}

// 				partLayout[hexString] = coords
// 			}
// 		}

// 		t, ok := pb.DynamicPartType_value[k]
// 		if !ok {
// 			err := fmt.Sprintf("invalid enum value, '%s'", k)
// 			return nil, errors.New(err)
// 		}

// 		m[pb.DynamicPartType(t)] = partLayout
// 	}

// 	return &BodyType{
// 		DisplayName: displayName,
// 		PartMaps:    m,
// 	}, nil
// }

// // StaticPart refers to a static image that should not transform. These images may move or rotate,
// // but they will not change mid-animation.
// type StaticPart struct {
// 	DisplayName        string
// 	BodyTypeIdentifier ID
// 	Type               *pb.StaticPartType
// 	Image              *image.Image
// 	Anchor             *Coordinates
// }

// // NewStaticPart creates a new StaticPart.
// func NewStaticPart(
// 	displayName string,
// 	bodyTypeIdentifier string,
// 	partType *pb.StaticPartType,
// 	image []byte,
// 	anchor *Coordinates,
// ) (*StaticPart, error) {
// 	r := bytes.NewReader(image)
// 	img, err := png.Decode(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &StaticPart{
// 		DisplayName:        displayName,
// 		BodyTypeIdentifier: bodyTypeIdentifier,
// 		Type:               partType,
// 		Image:              &img,
// 		Anchor:             anchor,
// 	}, nil
// }

// // DynamicPart refers to a static image that may transform mid-animation. These images do not
// // rotate or move, but they are represented in fluid, semi-3D space.
// type DynamicPart struct {
// 	DisplayName        string
// 	BodyTypeIdentifier ID
// 	Type               *pb.DynamicPartType
// 	PartLayout         map[Coordinates]string
// }

// // NewDynamicPart creates a new DynamicPart.
// func NewDynamicPart(
// 	displayName string,
// 	bodyTypeIdentifier ID,
// 	partType *pb.DynamicPartType,
// 	image []byte,
// ) (*DynamicPart, error) {
// 	m := make(map[Coordinates]string)

// 	r := bytes.NewReader(image)
// 	img, err := png.Decode(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	b := img.Bounds()
// 	for x := b.Min.X; x < b.Max.X; x++ {
// 		for y := b.Min.Y; y < b.Max.Y; y++ {
// 			r, g, b, a := b.At(x, y).RGBA()

// 			// Do not save transparent pixels.
// 			if a >= 255 {
// 				continue
// 			}

// 			hexString := fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
// 			coords := Coordinates{
// 				x: uint16(x),
// 				y: uint16(y),
// 			}

// 			m[coords] = hexString
// 		}
// 	}

// 	return &DynamicPart{
// 		DisplayName:        displayName,
// 		BodyTypeIdentifier: bodyTypeIdentifier,
// 		Type:               partType,
// 		PartLayout:         m,
// 	}, nil
// }

// // Prop refers to a static image that is body-type-agnostic.
// type Prop struct {
// 	DisplayName string
// 	Type        pb.PropType
// 	Image       *image.Image
// 	Anchor      *Coordinates
// }

// // NewProp creates a new Prop.
// func NewProp(
// 	displayName string,
// 	propType pb.PropType,
// 	image []byte,
// 	anchor *Coordinates,
// ) (*Prop, error) {
// 	r := bytes.NewReader(image)
// 	img, err := png.Decode(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Prop{
// 		DisplayName: displayName,
// 		Type:        propType,
// 		Image:       &img,
// 		Anchor:      anchor,
// 	}, nil
// }

// // Animation_Frame refers to a single frame of an animation. Animations are mostly a conglomerate of
// // these.
// type Animation_Frame struct {
// 	Layout map[Coordinates]string
// 	Parts  map[pb.StaticPartType]*Positioning
// 	Prop   *Positioning
// }

// // NewAnimationFrame creates a new Animation_Frame.
// func NewAnimationFrame(
// 	image []byte,
// 	parts map[string]*Positioning,
// 	prop *Positioning,
// ) (*Animation_Frame, error) {
// 	m := make(map[Coordinates]string)

// 	r := bytes.NewReader(image)
// 	img, err := png.Decode(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	b := img.Bounds()
// 	for x := b.Min.X; x < b.Max.X; x++ {
// 		for y := b.Min.Y; y < b.Max.Y; y++ {
// 			r, g, b, a := b.At(x, y).RGBA()

// 			// Do not save transparent pixels.
// 			if a >= 255 {
// 				continue
// 			}

// 			hexString := fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
// 			coords := Coordinates{
// 				x: uint16(x),
// 				y: uint16(y),
// 			}

// 			m[coords] = hexString
// 		}
// 	}

// 	p := make(map[pb.StaticPartType]*Positioning)
// 	for k, v := range parts {
// 		t, ok := pb.StaticPartType_value[k]
// 		if !ok {
// 			err := fmt.Sprintf("invalid enum value, '%s'", k)
// 			return nil, errors.New(err)
// 		}

// 		p[pb.StaticPartType(t)] = v
// 	}

// 	return &Animation_Frame{
// 		Layout: m,
// 		Parts:  p,
// 		Prop:   prop,
// 	}, nil
// }

// // Animation is the basis for any character's visible appearance. Still poses are just 1-frame
// // animations.
// type Animation struct {
// 	DisplayName        string
// 	BodyTypeIdentifier ID
// 	Allowed            []pb.PropType
// 	Frames             []*Animation_Frame
// }

// // NewAnimation creates a new instance of an Animation.
// func NewAnimation(
// 	displayName string,
// 	bodyTypeIdentifier string,
// 	allowed []string,
// 	frames []*Animation_Frame,
// ) (*Animation, error) {
// 	a := []pb.PropType{}

// 	for _, v := range allowed {
// 		t, ok := pb.PropType_value[v]
// 		if !ok {
// 			err := fmt.Sprintf("invalid enum value, '%s'", v)
// 			return nil, errors.New(err)
// 		}

// 		a = append(a, pb.PropType(t))
// 	}

// 	return &Animation{
// 		DisplayName:        displayName,
// 		BodyTypeIdentifier: bodyTypeIdentifier,
// 		Allowed:            a,
// 		Frames:             frames,
// 	}, nil
// }

// BodyTypesFilter is used to provide a limited list of the given item.
type BodyTypesFilter struct {
	BodyTypeIdentifiers []string
}

// StaticPartsFilter is used to provide a limited list of the given item.
type StaticPartsFilter struct {
	StaticPartTypes []pb.StaticPartType
}

// DynamicPartsFilter is used to provide a limited list of the given item.
type DynamicPartsFilter struct {
	DynamicPartTypes []pb.DynamicPartType
}

// PropsFilter is used to provide a limited list of the given item.
type PropsFilter struct {
	PropTypes []pb.PropType
}

// AnimationsFilter is used to provide a limited list of the given item.
type AnimationsFilter struct {
	BodyTypeIdentifiers []string
}

type CharacterImageDatabase interface {
	CreateBodyType(context.Context, *pb.BodyType) (ID, error)
	GetBodyType(context.Context, ID) (*pb.BodyType, error)
	ListBodyTypes(context.Context, *BodyTypesFilter) ([]pb.BodyType, error)

	CreateStaticPart(context.Context, *pb.StaticPart) (ID, error)
	GetStaticPart(context.Context, ID) (*pb.StaticPart, error)
	ListStaticParts(context.Context, *StaticPartsFilter) ([]pb.StaticPart, error)

	CreateDynamicPart(context.Context, *pb.DynamicPart) (ID, error)
	GetDynamicPart(context.Context, ID) (*pb.DynamicPart, error)
	ListDynamicParts(context.Context, *DynamicPartsFilter) ([]pb.DynamicPart, error)

	CreateProp(context.Context, *pb.Prop) (ID, error)
	GetProp(context.Context, ID) (*pb.Prop, error)
	ListProps(context.Context, *PropsFilter) ([]pb.Prop, error)

	CreateAnimation(context.Context, *pb.Animation) (ID, error)
	GetAnimation(context.Context, ID) (*pb.Animation, error)
	ListAnimations(context.Context, *AnimationsFilter) ([]pb.BodyType, error)
}
