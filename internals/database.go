package internals

import "fmt"

type DatabaseInterface interface {
	// New Initializes the database
	New(string) error
	// Store adds a new object to the database
	Store() error
}

type Database struct { 
	// path to db
	path string
}

func(db *Database) New(path string) error {
	db.path = path
	fmt.Println("New database create at", db.path)
	return nil
}

func (db *Database) Store() error {
	fmt.Println("Store called")
	return nil
}