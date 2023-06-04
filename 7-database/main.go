package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // this is the driver that we are going to use to connect with the database
	"github.com/google/uuid"
	// the underscore is called of blank identifier and is used to import a package that is not used in the code
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

// *Product -> is a pointer to a product because we want to change the value of the product for economy of memory
func NewProduct(name string, price float64) *Product {
	return &Product{
		// i want that to assign only once the id

		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	// to establish a connection with the database
	// the first parameter is the driver name
	// the second parameter is the data source name

	// the data source name is the username, password, host, port and the database name
	// the data source name is different for each database
	// the data source name for mysql is username:password@tcp(host:port)/database_name

	//<username>:<password>@tcp(<host>:<port>)/<database_name>

	// username: root
	// password: root
	// host: localhost
	// port: 3306
	// database name: goexpert

	// TCP -> is a network protocol that is used to establish a connection with the database of form secure and encrypted and garantize the integrity of the data that is sent and received and also garantize that the data is sent and received in the correct order without any data loss in the middle of the process on internet

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3305)/goexpert")
	if err != nil {
		panic(err.Error())
	}

	// close the connection with the database after the execution of the function
	defer db.Close()

	// create a new product

	product := NewProduct("soccer ball", 50.99)

	// insert the product in the database
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Name = "basketball ball"
	product.Price = 100.99

	// update the product in the database
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// cancel the context after the execution of the function with the defer statement to prevent memory leak and to prevent that the context is not used anymore after the execution of the function
	// ctx := context.WithValue(context.Background(), "id", product.ID)
	// ctx, cancel := context.WithCancel(ctx)
	// ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()

	// product, err = selectProduct(ctx, db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Product: %+v, have the price of: %.2f", product.Name, product.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Printf("Product: %+v, have the price of: %.2f\n", product.Name, product.Price)
	}

	err = deleteProduct(db, product.ID)
}

func insertProduct(db *sql.DB, product *Product) error {
	// the first parameter is the query
	// the second parameter is the arguments of the query
	// the arguments are used to prevent sql injection

	// concept of prepare statement -> is a query that is prepared to be executed multiple times with different arguments, protecting the database of sql injection

	// exist different types of sql injection, the most common is the sql injection that is used to delete the database or to get the data of the database

	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	// close the statement after the execution of the function
	defer stmt.Close()

	// execute the statement with the arguments
	_, err = stmt.Exec(product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}

	// close the statement after the execution of the function
	defer stmt.Close()

	// execute the statement with the arguments
	_, err = stmt.Exec(product.Name, product.Price, product.ID)

	if err != nil {
		return err
	}

	return nil
}

func selectProduct(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	// close the statement after the execution of the function
	defer stmt.Close()

	var product Product

	// search the product in the database and return only one row

	// &product.ID -> is the address of the id of the product, i am passing the address of the id of the product because i want to change the value of the id of the product, why? because i want to assign only once the id of the product in the memory and not every time that i search the product in the database and i want to change the value of the id of the product in the memory only once

	// its exactly the same order of the columns in the query -> id, name, price

	// with context is good practice to use the context in the query because if the query is taking a long time to execute, we can cancel the query with the context
	err = stmt.QueryRowContext(ctx, id).Scan(&product.ID, &product.Name, &product.Price)
	// err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]*Product, error) {
	// in this case i dont need to use the prepare statement because i dont have any arguments in the query and i dont need to protect the database of sql injection i can use directly the query method of the db object to execute the query in the database and return the rows of the query in the database

	// the query method of the db object is used to execute a query in the database and return the rows of the query in the database
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}

	// close the rows after the execution of the function
	defer rows.Close()

	// *Product -> is a pointer to a product because we want to change the value of the product for economy of memory
	var products []*Product

	// iterate over the rows of the query in the database
	for rows.Next() {
		var product Product

		// scan the row of the query in the database and assign the values of the row to the product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		// append the product to the products
		products = append(products, &product)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}

	// close the statement after the execution of the function
	defer stmt.Close()

	// execute the statement with the arguments

	// when i want execute a action in the database i use the exec method of the statement object and when i want to get the rows of the query in the database i use the query method of the statement object
	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
