package database

import (
	"log"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"gorm.io/gorm"
)

func migratePosts(db *gorm.DB, table interface{}) error {
	err := db.AutoMigrate(&table)
	if err != nil {
		log.Fatalf("failed to auto migrate table: %v", err)
	}

	err = db.Exec("ALTER TABLE posts ADD CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE").Error
	if err != nil {
		log.Fatalf("failed to add foreign key constraint: %v", err)
	}

	post := models.Post{
		ID:       1,
		Title:    "Hello, World!",
		Content:  "This is a test post.",
		AuthorID: 1,
	}

	err = db.Create(&post).Error
	if err != nil {
		log.Fatalf("failed to create post: %v", err)
	}

	return err
}
