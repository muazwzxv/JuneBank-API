package database

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

var db *GormInstance

type GormInstance struct {
	Config *Configuration
	Orm    *gorm.DB
}

type Configuration struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

func GetGormInstance() *GormInstance {
	if !db.isInstantiated() {
		config, err := readConfig()
		if err != nil {
			log.Fatalf("Error when reading config file %v", err)
		}

		db = newGorm(config)
	}
	return db
}

func newGorm(config *Configuration) *GormInstance {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DatabaseName,
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

func (g *GormInstance) isInstantiated() bool {
	return g.Orm != nil
}

func readConfig() (*Configuration, error) {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to open file %v", err)
	}

	data := make(map[string]Configuration)
	err = yaml.Unmarshal(file, data)

	if err != nil {
		return nil, err
	}

	config := data["database"]

	return &config, nil
}
