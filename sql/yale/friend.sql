select * from friend;

SELECT users.id,name 
FROM users 
JOIN friend
ON users.id = to_user
WHERE from_user = 1

INSERT INTO friend(from_user, to_user)
VALUES(
	( SELECT id FROM users WHERE name='nodev918'), ( SELECT id FROM users WHERE name='yale918') 
);

delete from friend where id= 6;



-- DROP TABLE friend;

-- CREATE TABLE "friend" (
--   "id" bigserial PRIMARY KEY NOT NULL,
--   "from_user" bigint NOT NULL,
--   "to_user" bigint NOT NULL
-- );


