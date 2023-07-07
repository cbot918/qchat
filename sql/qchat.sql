CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "email" varcahr NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestampz DEFAULT (now())
);

CREATE TABLE "friend" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "from" bigint NOT NULL,
  "to" bigint NOT NULL
);

ALTER TABLE "friend" ADD FOREIGN KEY ("friend_from") REFERENCES "users" ("id");

ALTER TABLE "friend" ADD FOREIGN KEY ("friend_to") REFERENCES "users" ("id");