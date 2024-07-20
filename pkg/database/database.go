package database

import (
	"log"

	"github.com/DanTDM2003/search-api-docker-redis/internal/appconfig/database"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"gorm.io/gorm"
)

func InitDatabase(p *database.PostgresConnection) error {
	err := createTables(p.DB)
	if err != nil {
		log.Fatalf("Could not create table: %v", err)
		return err
	}
	return nil
}

func tableExists(db *gorm.DB, model interface{}) bool {
	exists := db.Migrator().HasTable(&model)
	return exists
}

func createTables(db *gorm.DB) error {
	tables := []interface{}{models.MeteoriteLanding{}, models.User{}, models.Post{}}

	for _, table := range tables {
		var err error
		if !tableExists(db, table) {
			switch table.(type) {
			case models.MeteoriteLanding:
				err = migrateMetoriteLandings(db, table)
			case models.User:
				err = migrateUsers(db, table)
			case models.Post:
				err = migratePosts(db, table)
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}
