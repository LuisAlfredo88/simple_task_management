package sqlite

import (
	"log"
	securityEntity "stm/modules/security/domain/entity"
	"stm/shared/utility/encoding"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MigrationDBInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./cmd/bin/main.db"), &gorm.Config{})

	if err != nil {
		panic("Error connecting to the database")
	}

	return db
}

func loadInitialData(db *gorm.DB) {
	// Defining initial users
	users := []securityEntity.User{
		{
			Id:       "7414a010-b098-448c-8dd0-898c495bd9d6",
			Name:     "Admin",
			Password: encoding.GetMD5Hash("12345"),
			LastName: "Main",
			UserName: "admin",
			IsActive: 1,
		},
	}

	// Insert the initial users into the database
	for _, user := range users {
		result := db.FirstOrCreate(&user)
		if result.Error != nil {
			log.Println("Could not insert initial user:", result.Error)
		}
	}
}

func DoMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		// Security entities
		securityEntity.User{},
	)

	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	loadInitialData(db)
}
