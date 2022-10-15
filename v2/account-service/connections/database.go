package connections

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"junebank/v2/account-service/entity"
	"log"
)

var gormInstance = new(GormInstance)
var bunInstance = new(BunInstance)

type GormInstance struct {
	Config *Configuration
	Orm    *gorm.DB
}

type BunInstance struct {
	Config *Configuration
	Orm    *bun.DB
}

type Configuration struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func GetGormInstance() *GormInstance {
	if !gormInstance.isInstantiated() {
		config := readConfig()
		gormInstance = newGorm(config)
	}
	return gormInstance
}

func GetBunInstance() *BunInstance {
	if !bunInstance.isInstantiated() {
		config := readConfig()
		bunInstance = connectDB(config)
	}
	return bunInstance
}

func connectDB(config *Configuration) *BunInstance {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	bunOrm := bun.NewDB(conn, pgdialect.New())

	return &BunInstance{config, bunOrm}
}

func (g *GormInstance) Migrate() error {
	err := g.Orm.Debug().AutoMigrate(&entity.Account{})

	if err != nil {
		return err
	}

	return nil
}

func (g *GormInstance) isInstantiated() bool {
	return g.Orm != nil
}

func (b *BunInstance) isInstantiated() bool {
	return b.Orm != nil
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
		log.Fatalf("Error when connecting to connections %v", err)
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
		User:     reader.GetString("connections.user"),
		Host:     reader.GetString("connections.host"),
		Port:     reader.GetInt("connections.port"),
		Name:     reader.GetString("connections.name"),
		Password: reader.GetString("connections.password"),
	}
}
