// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).
// Code generated by sqlbrick. DO NOT EDIT IT.

// This file is generated from: pkg_info.sql

package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Type definition for PkgInfo which defined in sql file.
// This can be used as a model in database operation.
type PkgInfo struct {
	Id          int32  `db:"id"`
	Name        string `db:"name"`
	RootRepoUrl string `db:"root_repo_url"`
}

// Type definition for PkgInfoBrick. This brick will contains all database
// operation from given sql file. Each sql file will have only one brick.
type PkgInfoBrick struct {
	db *sqlx.DB
}

// newPkgInfoBrick will create a PkgInfo brick. This is used
// invoke the query function generated from sql file.
func newPkgInfoBrick(db *sqlx.DB) *PkgInfoBrick {
	return &PkgInfoBrick{db: db}
}

// CreateTable generated by sqlbrick, used to operate database table.
func (b *PkgInfoBrick) CreateTable() error {
	stmt, err := b.db.Prepare(`CREATE TABLE IF NOT EXISTS pkg_info (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  name text NOT NULL UNIQUE,
  root_repo_url text NOT NULL
)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

// InsertOne insert one record of package information into database.
func (b *PkgInfoBrick) InsertOne(args *PkgInfo) (sql.Result, error) {
	stmt, err := b.db.PrepareNamed(
		`INSERT INTO pkg_info (name, root_repo_url) VALUES (:name, :root_repo_url)`)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(args)
}

// DeleteById delete package information with given id.
func (b *PkgInfoBrick) DeleteById(id interface{}) (int64, error) {
	stmt, err := b.db.PrepareNamed(`DELETE FROM pkg_info WHERE id = :id`)
	if err != nil {
		return 0, err
	}

	// create map arguments for sqlx
	args := map[string]interface{}{
		"id": id,
	}

	result, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// UpdateById generated by sqlbrick, update data in database.
func (b *PkgInfoBrick) UpdateById(args *PkgInfo) (int64, error) {
	stmt, err := b.db.PrepareNamed(
		`UPDATE pkg_info SET root_repo_url = :root_repo_url WHERE id = :id`)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// SelectById query one record with given package name.
func (b *PkgInfoBrick) SelectById(dest interface{}, args interface{}) error {
	stmt, err := b.db.PrepareNamed(
		`SELECT id, (:host || '/' || name) as name, root_repo_url FROM pkg_info WHERE id = :id`)
	if err != nil {
		return err
	}

	row := stmt.QueryRowx(args)
	if row.Err() != nil {
		return row.Err()
	}

	return row.StructScan(dest)
}

// SelectByName query one record with given package name.
func (b *PkgInfoBrick) SelectByName(dest interface{}, args interface{}) error {
	stmt, err := b.db.PrepareNamed(
		`SELECT id, (:host || '/' || name) as name, root_repo_url FROM pkg_info
  WHERE name = :name`)
	if err != nil {
		return err
	}

	row := stmt.QueryRowx(args)
	if row.Err() != nil {
		return row.Err()
	}

	return row.StructScan(dest)
}

// SelectTotalPackages generated by sqlbrick, select data from database.
func (b *PkgInfoBrick) SelectTotalPackages(dest interface{}) error {
	stmt, err := b.db.Prepare(`SELECT COUNT (*) FROM pkg_info`)
	if err != nil {
		return err
	}
	return stmt.QueryRow().Scan(dest)
}

// SelectByPage query a list of records with given page.
func (b *PkgInfoBrick) SelectByPage(dest interface{}, args interface{}) error {
	stmt, err := b.db.PrepareNamed(
		`SELECT id, (:host || '/' || name) as name, root_repo_url FROM pkg_info LIMIT :limit OFFSET :offset`)
	if err != nil {
		return err
	}

	rows, err := stmt.Queryx(args)
	if err != nil {
		return err
	}

	return sqlx.StructScan(rows, dest)
}
