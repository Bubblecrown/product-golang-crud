package main

import (
	"database/sql"
	"fmt"
	"log"

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
	Id    int
	Name  string
	Price int
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

	// get product by id
	p, err := getAllProducts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get Product Successfully", p)

}

func createProduct(product *Product) error {
	_, err := db.Exec("INSERT INTO public.products(name, price) VALUES ($1, $2);", product.Name, product.Price)
	return err
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow(`SELECT id, name, price FROM products WHERE id = $1;`, id)
	err := row.Scan(&p.Id, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func getAllProducts() ([]Product, error) {
	var p []Product
	rows, err := db.Query(`SELECT id, name, price FROM products`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var productItem Product
		err := rows.Scan(&productItem.Id, &productItem.Name, &productItem.Price)
		if err != nil {
			return nil, err
		}
		p = append(p, productItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return p, nil
}

func updateProduct(id int, product *Product) (Product, error) {
	// exec รันคำสั่งแต่ไม่คืนค่าอะไรออกมา
	// QueryRow รันคำสั่ง และคืนค่าบางอย่างออกมา
	var p Product
	row := db.QueryRow("UPDATE public.products SET name=$1, price=$2 WHERE id=$3 RETURNING id, name, price;", product.Name, product.Price, id)
	err := row.Scan(&p.Id, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, err
}
