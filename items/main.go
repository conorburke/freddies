package main

import (
	"fmt"
	db "github.com/conorburke/freddies/items/database"
	"github.com/conorburke/freddies/items/models"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	initializeDb()

	registerRoutes(app)

	app.Listen(":3000")
}

func initializeDb() {
	var err error
	db.DB, err = gorm.Open(sqlite.Open("items.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	fmt.Println("Connected to sqlite3")
	// Migrate the schema
	db.DB.AutoMigrate(&models.Item{})

	fmt.Println("DB migrated successfully")

}

func registerRoutes(app *fiber.App) {
	baseURL := "api/items"
	app.Get(baseURL, helloWorld)
	app.Get(baseURL+"/second/:value", second)
	app.Post("/api/items", newItem)
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func second(c *fiber.Ctx) error {
	fmt.Fprintf(c, "Hello %s\n", c.Params("value"))
	return nil
}

func newItem(c *fiber.Ctx) error {
	db.DB.Create(&models.Item{Title: "Couch", OwnerID: "Kman", Category: "Junk"})
	var item models.Item
	db.DB.Last(&item)
	fmt.Println("item created!")
	fmt.Println(item)
	return c.JSON(item)
}
