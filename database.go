package main

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
	"fmt"
)

type DatabaseManager struct {
	db *sql.DB
}

func (this *DatabaseManager) Start () {
	var err error
	connection := os.Getenv("USER")+":"+os.Getenv("PASSWORD")+"@/"+os.Getenv("DBNAME")+"?charset=utf8"
	this.db, err = sql.Open("mysql", connection)

	if err != nil {
		panic(err)
	}

	configureDB(this.db);
}

func (this *DatabaseManager) GetItems (db *sql.DB) {
	rows, err := db.Query("SELECT * FROM items")

	if (err) != nil {
		panic(err)
	}

	fmt.Println(rows)
}

func configureDB (db *sql.DB) {
	createItemTable(db)
	createItemExampleTable(db)
}

func createItemTable (db *sql.DB) {
	_, err := db.Query("CREATE TABLE IF NOT EXISTS item (id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255), description VARCHAR(3000), price INT);")

	if err != nil {
		panic(err)
	}

	fmt.Println("Tabla items creada!")
}

func createItemExampleTable (db *sql.DB) {
	_, err := db.Query("CREATE TABLE IF NOT EXISTS item_example (item_id INT, img VARCHAR(500), PRIMARY KEY (item_id), FOREIGN KEY (item_id) REFERENCES item (id));")

	if err != nil {
		panic(err)
	}

	fmt.Println("Tabla items_example creada!")
}
