package models

import "gorm.io/gorm"

type Task struct {
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name      *string `json:"name"`
	Completed *bool   `json:"completed"`
}

func MigrateTasks(db *gorm.DB) error {
	err := db.AutoMigrate(&Task{})
	return err
}
