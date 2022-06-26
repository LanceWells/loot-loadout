CREATE TYPE dynamic_part_type AS ENUM (
  'LEFT_LEG',
  'RIGHT_LEG',
  'LEFT_ARM',
  'RIGHT_ARM'
);

CREATE TYPE static_part_type AS ENUM (
  'BODY',
  'HEAD',
  'HAIR'
);

CREATE TYPE prop_type AS ENUM (
  'WEAPON_MELEE_1H',
  'WEAPON_MELEE_2H',
  'WEAPON_RANGE_1H',
  'WEAPON_RANGE_2H'
);

CREATE TABLE color_string (
  id SERIAL PRIMARY KEY,
  hexstring varchar(9)
);

CREATE TABLE body_type (
  id SERIAL PRIMARY KEY,
  display_name varchar NOT NULL
);

CREATE TABLE dynamic_part (
  id SERIAL PRIMARY KEY,
  body_type_id int REFERENCES body_type(id),
  display_name varchar NOT NULL,
  part_type dynamic_part_type
);

CREATE TABLE dynamic_part_pixel (
  color_string_id int REFERENCES color_string,
  dynamic_part_id int REFERENCES dynamic_part(id),
  x smallint,
  y smallint,
  UNIQUE(x, y, dynamic_part_id),
  PRIMARY KEY(color_string_id, dynamic_part_id)
);

CREATE TABLE static_part (
  id SERIAL PRIMARY KEY,
  body_type_id int REFERENCES body_type(id),
  display_name varchar NOT NULL,
  part_type static_part_type
);

CREATE TABLE static_part_image (
  static_part_id int REFERENCES static_part(id) PRIMARY KEY,
  x smallint,
  y smallint,
  image_bytes bytea
);

CREATE TABLE prop (
  id SERIAL PRIMARY KEY,
  display_name varchar NOT NULL,
  part_type prop_type
);

CREATE TABLE prop_image (
  prop_id int REFERENCES prop(id) PRIMARY KEY,
  x smallint,
  y smallint,
  image_bytes bytea
);

CREATE TABLE animation (
  id SERIAL PRIMARY KEY,
  body_type_id int REFERENCES body_type(id),
  display_name varchar NOT NULL,
  part_type prop_type[]
);

CREATE TABLE animation_frame (
  id SERIAL PRIMARY KEY,
  animation_id int REFERENCES animation(id),
  frame_index int NOT NULL,
  UNIQUE(animation_id, frame_index)
);

CREATE TABLE animation_frame_pixel (
  color_string_id int REFERENCES color_string,
  animation_frame_id int REFERENCES animation_frame(id),
  x smallint,
  y smallint,
  UNIQUE(x, y, animation_frame_id),
  PRIMARY KEY(color_string_id, animation_frame_id)
);

CREATE TABLE animation_frame_static_position (
  animation_frame_id int REFERENCES animation_frame(id),
  part_type static_part_type,
  x smallint,
  y smallint,
  PRIMARY KEY(animation_frame_id, part_type)
);

CREATE TABLE animation_frame_prop_position (
  animation_frame_id int REFERENCES animation_frame(id),
  part_type prop_type,
  x smallint,
  y smallint,
  PRIMARY KEY(animation_frame_id, part_type)
);
