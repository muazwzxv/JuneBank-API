package database

import (
	"gopkg.in/yaml.v3"
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
	Post         int
	User         string
	Password     string
	databaseName string
}

func GetGormInstance() *GormInstance {
	if !db.isInstantiated() {
		config, err := readConfig()
		if err != nil {

		}

	}
	return nil
}

//func newGorm() *GormInstance {
//
//}

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
