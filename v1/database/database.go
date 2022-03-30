package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"junebank/entity"
	"log"
)

var db = new(GormInstance)

type GormInstance struct {
	Config *Configuration
	Orm    *gorm.DB
}

type Configuration struct {
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
	)

	if err != nil {
		return err
	}

	return nil
}

func (g *GormInstance) isInstantiated() bool {
	return g.Orm != nil
}

func newGorm(config *Configuration) *GormInstance {
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
		log.Fatalf("Error when connecting to database %v", err)
	}

	return &GormInstance{config, conn}
}

func readConfig() *Configuration {
	reader := viper.New()
	reader.SetConfigFile("config.yaml")

	if err := reader.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	return &Configuration{
		User:     reader.GetString("database.user"),
		Host:     reader.GetString("database.host"),
		Port:     reader.GetInt("database.port"),
		Name:     reader.GetString("database.name"),
		Password: reader.GetString("database.password"),
	}
}
