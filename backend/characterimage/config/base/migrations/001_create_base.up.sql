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

CREATE TABLE color_string (
  id SERIAL PRIMARY KEY,
  hexstring varchar(9)
);

CREATE TABLE body_type (
  id SERIAL PRIMARY KEY,
  display_name varchar NOT NULL
);

CREATE TABLE body_type_pixel (
  part_type dynamic_part_type,
  x smallint,
  y smallint,
  color_string_id int REFERENCES color_string,
  body_type_id int REFERENCES body_type(id),
  PRIMARY KEY(color_string_id, body_type_id),
  UNIQUE(x, y, body_type_id)
);

CREATE TABLE dynamic_part (
  id SERIAL PRIMARY KEY,
  display_name varchar NOT NULL,
  part_type dynamic_part_type,
  body_type_id int REFERENCES body_type(id)
);

CREATE TABLE dynamic_part_pixel (
  x smallint,
  y smallint,
  color_string_id int REFERENCES color_string,
  dynamic_part_id int REFERENCES dynamic_part(id),
  UNIQUE(x, y, dynamic_part_id),
  PRIMARY KEY(color_string_id, dynamic_part_id)
);

CREATE TABLE static_part (
  id SERIAL PRIMARY KEY,
  display_name varchar NOT NULL,
  part_type static_part_type,
  image_bytes bytea,
  x smallint,
  y smallint,
  body_type_id int REFERENCES body_type(id)
);

CREATE TABLE dynamic_part_colormap_pixel (
  part_type dynamic_part_type,
  x smallint,
  y smallint,
  color_string_id int REFERENCES color_string,
  dynamic_part_id int REFERENCES dynamic_part(id),
  PRIMARY KEY(color_string_id, dynamic_part_id),
  UNIQUE(x, y, dynamic_part_id)
);
