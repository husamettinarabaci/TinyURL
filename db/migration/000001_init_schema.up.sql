CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
  "url_count" bigint NOT NULL,
  "user_type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transforms" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "long_url" varchar NOT NULL,
  "short_url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "transforms" ("owner");

CREATE UNIQUE INDEX ON "transforms" ("owner", "long_url");

ALTER TABLE "transforms" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");