CREATE TABLE if not exists "room"
(
  "id"                  uuid not null PRIMARY KEY,
  "room_num"            integer,
  "stage"               integer,
  "last_cleaned"        timestamp,
  "ocupied"             boolean
);

CREATE TABLE if not exists "guest" 
(
  "id"                  uuid not null PRIMARY KEY,
  "first_name"          varchar,
  "second_name"         varchar,
  "current_room_num"    uuid,
  "created_at"          timestamp,
  "checked_in"          timestamp,
  "checked_out"         timestamp,
  "deleted_at"          timestamp
);

COMMENT ON COLUMN "room"."ocupied" IS 'Ocupied/Unocupied';

ALTER TABLE "guest" ADD FOREIGN KEY ("current_room_num") REFERENCES "room" ("id");


INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 101, 1, CURRENT_TIMESTAMP, false);
INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 102, 1, CURRENT_TIMESTAMP, false);
INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 103, 1, CURRENT_TIMESTAMP, false);
INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 104, 1, CURRENT_TIMESTAMP, false);
INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 105, 1, CURRENT_TIMESTAMP, false);
INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 106, 1, CURRENT_TIMESTAMP, false);
INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 201, 2, CURRENT_TIMESTAMP, false);
INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 202, 2, CURRENT_TIMESTAMP, false);
INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 203, 2, CURRENT_TIMESTAMP, false);
INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 204, 2, CURRENT_TIMESTAMP, false);
INSERT INTO room (id, room_num, stage, last_cleaned, ocupied) VALUES (uuid_generate_v4(), 205, 2, CURRENT_TIMESTAMP, false);