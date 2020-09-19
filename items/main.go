package main

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/conorburke/freddies/items/models"
	"github.com/conorburke/freddies/items/database"
)

func main() {
	app := fiber.New();
	app.Use(cors.New())

	initializeDb()
	defer database.DBConn.Close()

	registerRoutes(app)

	app.Listen(":3000")
}

func initializeDb() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "item.db")
	if err != nil {
		// panic("Failed to connect to DB")
		fmt.Println(err)
	}
	fmt.Println("Connected to sqlite3")
	database.DBConn.AutoMigrate(&models.Item{})
	fmt.Println("DB migrated successfully")

}

func registerRoutes(app *fiber.App) {
	baseURL := "api/items"
	app.Get(baseURL, helloWorld)
	app.Get(baseURL + "/second/:value", second)
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
	db := database.DBConn
	var item models.Item
	item.Title = "Couch"
	item.Owner = "Kman"
	item.Type = "Junk"
	db.Create(&item)
	fmt.Println("item created!")
	return c.JSON(item)
}


