package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // this is the driver that we are going to use to connect with the database
	// the underscore is called of blank identifier and is used to import a package that is not used in the code

	"github.com/google/uuid"
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

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err.Error())
	}

	// close the connection with the database after the execution of the function
	defer db.Close()

	// create a new product

	product := NewProduct("Notebook", 10000.99)

	// insert the product in the database
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	// the first parameter is the query
	// the second parameter is the arguments of the query
	// the arguments are used to prevent sql injection

	// concept of prepareStatement -> is a query that is prepared to be executed multiple times with different arguments, protecting the database of sql injection

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
