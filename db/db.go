package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"skillspar/user_service/model"
	"skillspar/user_service/store/impl"
	"skillspar/user_service/utils"
)

var DbConfig = func() *mysql.Config {
	return &mysql.Config{
		User:                 utils.Getenv("DB_USERNAME", "root"),
		Passwd:               utils.Getenv("DB_PASSWORD", "root"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", utils.Getenv("DB_HOST", "localhost"), utils.Getenv("DB_PORT", "3306")),
		DBName:               utils.Getenv("DB_DATABASE", "db"),
		ParseTime:            true,
		AllowNativePasswords: true,
	}
}()

func New() *gorm.DB {

	//dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	DbConfig.User, DbConfig.Passwd, DbConfig.Net, DbConfig.Addr, DbConfig.DBName)
	//
	//fmt.Println(dsn)

	db, dbErr := sql.Open("mysql", DbConfig.FormatDSN())
	if dbErr != nil {
		log.Fatalf("Can't connect to database %v with user: %v, pass: %v, dbname: %v\n",
			DbConfig.Addr, DbConfig.User, DbConfig.Passwd, DbConfig.DBName)
	}

	gormDB, gormDbErr := gorm.Open(gormMysql.New(gormMysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if gormDbErr != nil {
		log.Fatalln(gormDbErr)
	}

	return gormDB
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Client{},
		&model.PersonalAccessClient{},
		&model.AccessToken{},
		&model.RefreshToken{},
	)
}

func Seeding(db *gorm.DB) {
	client := &model.Client{
		UserId:               1,
		Name:                 "SkillSpar",
		Secret:               "zxcvcvxvcxvxcvzcx",
		Provider:             "thangnmm",
		Redirect:             "skillspar.com",
		PersonalAccessClient: false,
		Revoked:              false,
	}

	cs := impl.NewClientStore(db)

	cs.Create(client)

}
