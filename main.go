package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// ตัวแปลสำหรับการ connection
const (
	host         = "0.0.0.0" // ตามชื่อ container_name
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

// global variables
var db *sql.DB

// database struct
type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	// connect to database
	tempdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = tempdb
	defer db.Close()
	// Check the connection
	// Ping to database and database response back
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")
	// create
	// err = createProduct(&Product{Name: "Go-product 1", Price: 389})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Create Product Successfully")

	// update
	// productUpdate, err := updateProduct(9, &Product{Name: "Thirsd novel", Price: 769})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Update Product Successfully", productUpdate)

	// get product by id
	// product, err := getProduct(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Get Product Successfully", product)

	// get all product
	// p, err := getAllProducts()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Get Product Successfully", p)
	app := fiber.New()
	app.Get("/product/:id", getProductsByIdHandler)
	app.Get("/products", getAllProductsHandler)
	app.Post("/create", createProductHandler)
	app.Post("/product/:id/edit", updateProductHandler)
	app.Delete("/product/:id", deleteProductHandler)
	app.Listen(":8080")
}

func getAllProductsHandler(c *fiber.Ctx) error {
	product, err := getAllProducts()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func getProductsByIdHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	product, err := getProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func createProductHandler(c *fiber.Ctx) error {
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := createProduct(p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(p)
}

func updateProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := updateProduct(productId, p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func deleteProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = deleteProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusOK)
}
