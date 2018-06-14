{define name CreateTable}
CREATE TABLE IF NOT EXISTS user (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  name text NOT NULL UNIQUE,
  password text NOT NULL,
  salt text NOT NULL,
  role int NOT NULL
);
{end define}

-- insert one record of user into database.
{define name InsertOne}
INSERT INTO user (name, password, salt, role) VALUES (${name}, ${password}, ${salt}, ${role});
{end define}

-- delete on recod of user in database.
{define name DeleteById}
DELETE FROM user WHERE id = ${id};
{end define}

-- update user's information.
{define name UpdateById}
UPDATE user SET password = ${password} WHERE id = ${id};
{end define}

-- query one record with given username.
{define name SelectByName, mapper single}
SELECT * FROM user WHERE name = ${name} LIMIT 1;
{end define}

-- query one record with given username.
{define name SelectById, mapper single}
SELECT * FROM user WHERE id = ${id} LIMIT 1;
{end define}

