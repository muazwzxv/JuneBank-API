package database

import (
	"fmt"
	"junebank_v1/entity"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db = new(GormInstance)

type GormInstance struct {
	Config *DatabaseConfiguration
	Orm    *gorm.DB
}

type DatabaseConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func GetGormInstance() *GormInstance {
	if !db.isInstantiated() {
		config := readConfig()
		db = newGorm(config)
	}
	return db
}

func (g *GormInstance) Migrate() error {
	err := g.Orm.Debug().AutoMigrate(
		&entity.Account{},
		&entity.Transaction{},
	)

	if err != nil {
		return err
	}

	return nil
}

func (g *GormInstance) isInstantiated() bool {
	return g.Orm != nil
}

func newGorm(config *DatabaseConfiguration) *GormInstance {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatalf("Error when connecting to connections %v", err)
	}

	return &GormInstance{config, conn}
}

func readConfig() *DatabaseConfiguration {
	reader := viper.New()
	reader.SetConfigFile("config.yaml")

	if err := reader.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	return &DatabaseConfiguration{
		User:     reader.GetString("connections.user"),
		Host:     reader.GetString("connections.host"),
		Port:     reader.GetInt("connections.port"),
		Name:     reader.GetString("connections.name"),
		Password: reader.GetString("connections.password"),
	}
}
