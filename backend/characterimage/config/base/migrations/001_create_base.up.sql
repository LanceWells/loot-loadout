-- Dynamic parts are represented as an image comprised of differently-colored pixels that map to an
-- animation frame's set of differently-colored pixels. This gives us a lookup equal to:
--
-- frame_coord => frame_color => mapping_color => mapping_coord => part_coord => part_color
-- \_________/                                                                   \________/
--      \\__________________________________________________________________________//
--
--    1. Get all of the pixels of an animation frame.
--    2. JOIN the pixels of the dynamic_part_mapping(s) for the appropriate body_type_id.
--    3. JOIN the pixels of the dynamic_part(s) for the appropriate id(s).
CREATE TYPE dynamic_part_type AS ENUM (
  'LEFT_LEG',
  'RIGHT_LEG',
  'LEFT_ARM',
  'RIGHT_ARM'
);

-- Static parts are represented as .png images that do not change (plan to add color-swap support in
-- future, but no transformation).
CREATE TYPE static_part_type AS ENUM (
  'BODY',
  'HEAD',
  'HAIR',
  'EYES',
  'ACCESSORY_HEAD'
);

-- Props are nearly the same as static parts, but they differ in that they are not restricted by a
-- body type, but are instead restricted by an animation's allowed prop types.
CREATE TYPE prop_type AS ENUM (
  'WEAPON_MELEE_1H',
  'WEAPON_MELEE_2H',
  'WEAPON_RANGE_1H',
  'WEAPON_RANGE_2H'
);

-- expression_type describes the expression for a character. This is defined per-animation-frame.
CREATE TYPE expression_type AS ENUM (
  'NEUTRAL'
);

-- color_string is a means to normalize color lookups (any given color should show up at least
-- twice).
CREATE TABLE color_string (
  id        SERIAL PRIMARY KEY,
  hexstring VARCHAR(9) NOT NULL,
  UNIQUE(hexstring)
);

-- body_type is the basis for all other character display information.
CREATE TABLE body_type (
  id            SERIAL PRIMARY KEY,
  display_name  VARCHAR NOT NULL,
  height        SMALLINT NOT NULL,
  width         SMALLINT NOT NULL
);

-- dynamic_part_mapping is represented as a single image that represents a given body part rotated
-- in 3D space. We save this image as the color value at each pixel, as well as the mapping\
-- information for this part.
CREATE TABLE dynamic_part_mapping (
  ID            SERIAL PRIMARY KEY,
  body_type_id  INT NOT NULL REFERENCES body_type(id) ON DELETE CASCADE,
  part_type     dynamic_part_type NOT NULL,
  UNIQUE(body_type_id, part_type)
);

CREATE TABLE dynamic_part_mapping_pixel (
  color_string_id         INT NOT NULL REFERENCES color_string(id) ON DELETE CASCADE,
  dynamic_part_mapping_id INT NOT NULL REFERENCES dynamic_part_mapping(id) ON DELETE CASCADE,
  x                       SMALLINT NOT NULL,
  y                       SMALLINT NOT NULL,
  UNIQUE(x, y, dynamic_part_mapping_id),
  PRIMARY KEY(color_string_id, dynamic_part_mapping_id)
);

-- dynamic_part is represented by an image of what a given body part should look like after it has
-- been applied to a character directly. The position of each pixel is referenced by
-- dynamic_part_mapping, which gives us our pixel-to-color lookup.
CREATE TABLE dynamic_part (
  id                      SERIAL PRIMARY KEY,
  dynamic_part_mapping_id INT NOT NULL REFERENCES dynamic_part_mapping(id) ON DELETE CASCADE,
  display_name            VARCHAR NOT NULL NOT NULL,
  part_type               dynamic_part_type NOT NULL
);

CREATE TABLE dynamic_part_pixel (
  id              SERIAL PRIMARY KEY,
  color_string_id INT NOT NULL REFERENCES color_string(id) ON DELETE CASCADE,
  dynamic_part_id INT NOT NULL REFERENCES dynamic_part(id) ON DELETE CASCADE,
  x               SMALLINT NOT NULL,
  y               SMALLINT NOT NULL,
  UNIQUE(x, y, dynamic_part_id)
);

CREATE TABLE dynamic_part_thumbnail (
  dynamic_part_id   INT REFERENCES dynamic_part(id) ON DELETE CASCADE PRIMARY KEY,
  image_bytes       bytea NOT NULL
);

-- static_part is represented by a static image that does not transform (but might rotate). Any
-- static part may only be applied to a single body type.
CREATE TABLE static_part (
  id SERIAL PRIMARY KEY,
  body_type_id  INT NOT NULL REFERENCES body_type(id) ON DELETE CASCADE,
  display_name  VARCHAR NOT NULL,
  part_type     static_part_type NOT NULL
);

CREATE TABLE static_part_image (
  static_part_id  INT REFERENCES static_part(id) ON DELETE CASCADE PRIMARY KEY,
  x               SMALLINT NOT NULL,
  y               SMALLINT NOT NULL,
  image_bytes     bytea NOT NULL
);

-- prop is represented by a static image that does not transform (but might rotate). These items may
-- be applied to any body type, but are limited by the current animation for a character.
CREATE TABLE prop (
  id            SERIAL PRIMARY KEY,
  display_name  VARCHAR NOT NULL,
  part_type     prop_type NOT NULL
);

CREATE TABLE prop_image (
  prop_id     INT REFERENCES prop(id) ON DELETE CASCADE PRIMARY KEY,
  x           SMALLINT NOT NULL,
  y           SMALLINT NOT NULL,
  image_bytes bytea NOT NULL
);

-- animation is the representational information for a character's pose and how its body is laid-out
-- visually.
CREATE TABLE animation (
  id            SERIAL PRIMARY KEY,
  body_type_id  INT NOT NULL REFERENCES body_type(id) ON DELETE CASCADE,
  display_name  VARCHAR NOT NULL NOT NULL,
  part_type     prop_type[] NOT NULL
);

-- animation_frame is a single frame of animation for a larger animation. In concert, these make a
-- series of motions for a character. Note that each frame includes positional data, but it also
-- includes the basic dynamic_part_mapping lookup colors that we use to determine the final output.
CREATE TABLE animation_frame (
  id            SERIAL PRIMARY KEY,
  animation_id  INT NOT NULL REFERENCES animation(id) ON DELETE CASCADE,
  frame_index   INT NOT NULL,
  expression    expression_type NOT NULL,
  duration      SMALLINT NOT NULL,
  UNIQUE(animation_id, frame_index)
);

CREATE TABLE animation_frame_pixel (
  color_string_id     INT REFERENCES color_string NOT NULL,
  animation_frame_id  INT NOT NULL REFERENCES animation_frame(id) ON DELETE CASCADE,
  x                   SMALLINT NOT NULL,
  y                   SMALLINT NOT NULL,
  UNIQUE(x, y, animation_frame_id),
  PRIMARY KEY(color_string_id, animation_frame_id)
);

CREATE TABLE animation_frame_static_position (
  animation_frame_id  INT NOT NULL REFERENCES animation_frame(id) ON DELETE CASCADE,
  part_type           static_part_type NOT NULL,
  x                   SMALLINT NOT NULL,
  y                   SMALLINT NOT NULL,
  rotation            SMALLINT NOT NULL,
  PRIMARY KEY(animation_frame_id, part_type)
);

CREATE TABLE animation_frame_prop_position (
  animation_frame_id  INT REFERENCES animation_frame(id) ON DELETE CASCADE PRIMARY KEY,
  x                   SMALLINT NOT NULL,
  y                   SMALLINT NOT NULL,
  rotation            SMALLINT NOT NULL
);
