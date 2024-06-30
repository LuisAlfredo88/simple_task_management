package sqlite

import (
	"log"
	securityEntity "stm/modules/security/domain/entity"
	taskManagementEntity "stm/modules/task_management/domain/entity"
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

func DoMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		// Security entities
		securityEntity.User{},

		// Task management entities
		taskManagementEntity.TaskStatus{},
		taskManagementEntity.Task{},
	)

	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	loadInitialData(db)
}

func loadInitialData(db *gorm.DB) {
	// Defining initial users
	users := []securityEntity.User{
		{
			Id:       "7414a010-b098-448c-8dd0-898c495bd9d6",
			Name:     "Main",
			Password: encoding.GetMD5Hash("12345"),
			LastName: "First",
			UserName: "main",
			IsActive: 1,
		},

		{
			Id:       "448c0009-b098-448c-8dd0-7414a0104457",
			Name:     "Guest",
			Password: encoding.GetMD5Hash("12345"),
			LastName: "Second",
			UserName: "guest",
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

	// Defining initial task status
	taskStatus := []taskManagementEntity.TaskStatus{
		{Id: 1, Name: "to_do", Color: "#03A9F4"},
		{Id: 2, Name: "in_progress", Color: "#FF9800"},
		{Id: 3, Name: "done", Color: "#21b421"},
	}

	// Insert the initial task status into the database
	for _, status := range taskStatus {
		result := db.FirstOrCreate(&status)
		if result.Error != nil {
			log.Println("Could not insert initial task status:", result.Error)
		}
	}

	// Defining initial task status
	tasks := []taskManagementEntity.Task{
		{
			Id:          1,
			Title:       "Task #1 from main user",
			Description: "This is the first registered task",
			StatusId:    1,
			CreatedById: "7414a010-b098-448c-8dd0-898c495bd9d6",
		},

		{
			Id:          2,
			Title:       "Task #2 from main user",
			Description: "This is the second registered task",
			StatusId:    2,
			CreatedById: "7414a010-b098-448c-8dd0-898c495bd9d6",
		},

		{
			Id:           3,
			Title:        "Task #3 from main user",
			Description:  "This is the third registered task, assigned to the guest user",
			StatusId:     2,
			CreatedById:  "7414a010-b098-448c-8dd0-898c495bd9d6",
			AssignedToId: "448c0009-b098-448c-8dd0-7414a0104457",
		},
	}

	// Insert the initial task status into the database
	for _, task := range tasks {
		result := db.FirstOrCreate(&task)
		if result.Error != nil {
			log.Println("Could not insert initial task:", result.Error)
		}
	}

}
