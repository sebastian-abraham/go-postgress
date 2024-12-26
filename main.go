package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/mod/sumdb/storage"
	"gorm.io/gorm"
)

type Task struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetTasks(context *fiber.Ctx) error {
	taskModels := &[]models.Task{}

	err := r.DB.Find(taskModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "could not get the tasks",
			})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "tasks retrieved successfully",
			"data":    taskModels,
		})

	return nil
}

func (r *Repository) CreateTask(context *fiber.Ctx) error {
	task := Task{}

	err := context.BodyParser(task)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "request Failed",
			})

		return err
	}

	err = r.DB.Create(&task).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Could not create book",
			})

		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "task created successfully",
		})

	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/task", r.CreateTask)
	api.Delete("/task/:id", r.DeleteTask)
	api.Get("/task/:id", r.GetTask)
	api.Get("/task", r.GetTasks)

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load the database")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
