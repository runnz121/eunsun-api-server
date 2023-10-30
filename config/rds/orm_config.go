package rds

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

var (
	DB      *gorm.DB
	err     error
	DBError error
)

func SetUpGorm() error {
	var db = DB

	masterConnection, slaveConnection := RDSConfiguration()

	debug := false

	db, err = gorm.Open(mysql.Open(masterConnection), &gorm.Config{
		QueryFields:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if !debug {
		err = db.Use(
			dbresolver.Register(dbresolver.Config{
				Replicas: []gorm.Dialector{
					mysql.Open(slaveConnection),
				},
				Policy: dbresolver.RandomPolicy{},
			}).
				SetConnMaxIdleTime(5 * time.Minute).
				SetConnMaxLifetime(5 * time.Minute).
				SetMaxIdleConns(1).
				SetMaxOpenConns(1),
		)
	}

	if err != nil {
		DBError = err
		return err
	}

	DB = db
	return nil
}

func Gorm() *gorm.DB {
	return DB
}
