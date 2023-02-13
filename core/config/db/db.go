package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/itokun99/ms-go-boilerplate/core/model"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDB(conf model.ServerConfig) *Database {
	var DB *gorm.DB
	var err error

	var host, user, password, name, port string

	defer func() {
		if r := recover(); r != nil {
			log.Println("core/config/db", errors.New("recover"), r)
		}
	}()

	// check DB version
	host = conf.DBConfig.Host
	port = conf.DBConfig.Port
	user = conf.DBConfig.User
	password = conf.DBConfig.Password
	name = conf.DBConfig.Name

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, name, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err, "core/config/db: gorm open connect")
	}

	var ctx context.Context

	DB = DB.WithContext(ctx)

	dbSQL, err := DB.DB()
	if err != nil {
		log.Fatalln(err, "core/config/db: gorm open connect", nil)
	}

	//Database Connection Pool
	dbSQL.SetMaxIdleConns(10)
	dbSQL.SetMaxOpenConns(100)
	dbSQL.SetConnMaxLifetime(time.Hour)

	err = dbSQL.Ping()
	if err != nil {
		log.Fatal(err, "config/DB: can't ping the DB, WTF", nil)
	} else {
		go doEvery(10*time.Minute, pingDb, DB)
		return &Database{
			DB: DB,
		}
	}

	return &Database{
		DB: DB,
	}
}

func doEvery(d time.Duration, f func(*gorm.DB), x *gorm.DB) {
	for range time.Tick(d) {
		f(x)
	}
}

func pingDb(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		log.Println(err, "core/config/db func pingDB: recover from contract db init")
	}

	err = dbSQL.Ping()
	if err != nil {
		log.Println(err, "core/config/db func pingDB: recover from contract db init")
	}
}

func (d *Database) AutoMigrate(schemas ...interface{}) {
	for _, schema := range schemas {
		if err := d.DB.AutoMigrate(schema); err != nil {
			log.Println(err, "core/config/db func AutoMigrate: recover from contract db init")
		}
	}
}

func (db *Database) DropTable(schemas ...interface{}) error {
	for _, schema := range schemas {

		if err := db.DB.Migrator().DropTable(schema); err != nil {
			log.Println(err, "Config/db/driver/ func DropTable: recover from contract db init")
			return err
		}
	}
	return nil
}
