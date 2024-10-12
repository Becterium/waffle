package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"waffle/app/user/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMysqlClient, NewUserRepo, NewAddressRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, mysqlClient *gorm.DB) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data"))
	data := &Data{
		db: mysqlClient,
	}
	db, _ := data.db.DB()
	return data, func() {
		if err := db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func NewMysqlClient(c *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/db"))

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	if err = db.AutoMigrate(); err != nil {
		log.Fatal("failed to Database AutoMigrate")
	}
	return db
}
