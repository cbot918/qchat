CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "friend" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "friend_from" bigint NOT NULL,
  "friend_to" bigint NOT NULL
);

ALTER TABLE "friend" ADD FOREIGN KEY ("friend_from") REFERENCES "users" ("id");

ALTER TABLE "friend" ADD FOREIGN KEY ("friend_to") REFERENCES "users" ("id");

INSERT INTO users(name, email,password) values ('yale918','yale918@gmail.com','12345');
INSERT INTO users(name, email,password) values ('nodev918','nodev918@gmail.com','12345');
INSERT INTO users(name, email,password) values ('kitty','kitty@gmail.com','12345');
INSERT INTO users(name, email,password) values ('nami','nami@gmail.com','12345');
INSERT INTO users(name, email,password) values ('dora','dora@gmail.com','12345');
INSERT INTO users(name, email,password) values ('jojo','jojo@gmail.com','12345');

DELETE FROM users where id=5;

INSERT INTO friend ()

select * from users; 