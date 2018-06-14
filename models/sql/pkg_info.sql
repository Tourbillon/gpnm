{define name CreateTable}
CREATE TABLE IF NOT EXISTS pkg_info (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  name text NOT NULL UNIQUE,
  root_repo_url text NOT NULL
);
{end define}

-- insert one record of package information into database.
{define name InsertOne}
INSERT INTO pkg_info (name, root_repo_url) VALUES (${name}, ${root_repo_url});
{end define}

-- delete package information with given id.
{define name DeleteById}
DELETE FROM pkg_info WHERE id = ${id};
{end define}

{define name UpdateById}
UPDATE pkg_info SET root_repo_url = ${root_repo_url} WHERE id = ${id};
{end define}

-- query one record with given package name.
{define name SelectById, mapper single}
SELECT id, (${host} || '/' || name) as name, root_repo_url FROM pkg_info WHERE id = ${id};
{end define}

-- query one record with given package name.
{define name SelectByName, mapper single}
SELECT id, (${host} || '/' || name) as name, root_repo_url FROM pkg_info
  WHERE name = ${name};
{end define}

{define name SelectTotalPackages, mapper single}
SELECT COUNT (*) FROM pkg_info;
{end define}

-- query a list of records with given page.
{define name SelectByPage}
SELECT id, (${host} || '/' || name) as name, root_repo_url FROM pkg_info LIMIT ${limit} OFFSET ${offset}
{end define}