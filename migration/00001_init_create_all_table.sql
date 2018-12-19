-- +goose Up
-- SQL in this section is executed when the migration is applied.:wqw

CREATE TABLE "class" (
  "id" uuid,
  "name" varchar(255),
  "category_id" uuid,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "category" (
  "id" uuid,
  "name" varchar(255),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "address" (
  "id" uuid,
  "name" varchar(255),
  "street1" varchar(255),
  "street2" varchar(255),
  "city" varchar(255),
  "state" varchar(255),
  "country" varchar(255),
  "zip_code" varchar(255),
  "lat" double precision,
  "lng" double precision,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "center" (
  "id" uuid,
  "name" varchar(255),
  "address_id" uuid,
  "city_id" uuid,
  "open_at" varchar(100),
  "close_at" varchar(100),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "schedule" (
  "id" uuid,
  "center_id" uuid,
  "start_at" TIMESTAMPTZ,
  "end_at" TIMESTAMPTZ,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "schedule_detail" (
  "id" uuid,
  "schedule_id" uuid,
  "class_id" uuid,
  "trainer_id" uuid,
  "start_at" TIMESTAMPTZ,
  "end_at" TIMESTAMPTZ,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "trainer" (
  "id" uuid,
  "name" varchar(255),
  "image_id" uuid,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "user" (
  "id" uuid,
  "name" varchar(255),
  "fb_id" varchar(255),
  "image_id" uuid,
  "email" varchar(255),
  "identity" varchar(255),
  "encrypt_password" varchar(255),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "role" (
  "id" uuid,
  "name" varchar(255),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "role_permission" (
  "id" uuid,
  "role_id" uuid,
  "permission_id" uuid,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "permission" (
  "id" uuid,
  "name" varchar(255),
  "model" varchar(255),
  "action" varchar(255),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "user_role" (
  "id" uuid,
  "role_id" uuid,
  "user_id" uuid,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "user_session" (
  "id" uuid,
  "user_id" uuid,
  "token" varchar(255),
  "expired_at" TIMESTAMPTZ,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "tag" (
  "id" uuid,
  "name" varchar(255),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "tag_map" (
  "id" uuid,
  "tag_id" uuid,
  "resource" varchar(255),
  "resource_id" uuid,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "image" (
  "id" uuid,
  "url" varchar(255),
  "source" varchar(255),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "image_map" (
  "id" uuid,
  "image_id" uuid,
  "resource" varchar(255),
  "resource_id" uuid,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

CREATE TABLE "city" (
  "id" uuid,
  "name" varchar(255),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id")
);

ALTER TABLE "image_map" ADD FOREIGN KEY ("image_id") REFERENCES "image" ("id");

ALTER TABLE "tag_map" ADD FOREIGN KEY ("tag_id") REFERENCES "tag" ("id");

ALTER TABLE "user_role" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_role" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("image_id") REFERENCES "image" ("id");

ALTER TABLE "trainer" ADD FOREIGN KEY ("image_id") REFERENCES "image" ("id");

ALTER TABLE "user_session" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "center" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");

ALTER TABLE "center" ADD FOREIGN KEY ("city_id") REFERENCES "city" ("id");

ALTER TABLE "schedule" ADD FOREIGN KEY ("center_id") REFERENCES "center" ("id");

ALTER TABLE "schedule_detail" ADD FOREIGN KEY ("schedule_id") REFERENCES "schedule" ("id");

ALTER TABLE "schedule_detail" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");

ALTER TABLE "schedule_detail" ADD FOREIGN KEY ("trainer_id") REFERENCES "trainer" ("id");

ALTER TABLE "class" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
