// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package models

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"

	"anbillon.com/gpnm/modules/setting"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/pbkdf2"
)

var (
	sqlbrick *SqlBrick
)

const (
	RoleAdmin  = 1
	RoleNormal = 2
)

//go:generate sqlbrick -w ./sql -o ./
func init() {
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", setting.Config.DbPath)
	db, err := sqlx.Connect("sqlite3", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	sqlbrick = NewSqlBrick(db)
	if err := sqlbrick.PkgInfo.CreateTable(); err != nil {
		log.Fatalln(err)
	}
	if err := sqlbrick.User.CreateTable(); err != nil {
		log.Fatalln(err)
	}
	u := &User{
		Name:     "admin",
		Password: "admin",
		Salt:     GenRandString(10),
		Role:     RoleAdmin,
	}
	u.EncryptPassword()
	sqlbrick.User.InsertOne(u)
}

// GetSqlBrick will return SqlBrick instance.
func GetSqlBrick() *SqlBrick {
	return sqlbrick
}

// GenRandString will generate random string with number, lower and upper letters.
func GenRandString(n int, alphabets ...byte) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		if len(alphabets) == 0 {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(bytes)
}

// EncryptPassword encrypt password with salt.
func (u *User) EncryptPassword() {
	key := pbkdf2.Key([]byte(u.Password), []byte(u.Salt), 4096, 32, sha256.New)
	u.Password = fmt.Sprintf("%x", key)
}

// ValidatePassword validate user's password.
func (u *User) ValidatePassword(password string) bool {
	user := User{
		Password: password,
		Salt:     u.Salt,
	}
	user.EncryptPassword()

	return subtle.ConstantTimeCompare([]byte(u.Password), []byte(user.Password)) == 1
}
