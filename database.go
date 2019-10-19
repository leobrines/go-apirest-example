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
func (this *DatabaseManager) start () {
	var err error
	connection := os.Getenv("DBUSER")+":"+os.Getenv("DBPASSWD")+"@/"+os.Getenv("DBNAME")+"?charset=utf8"
	this.db, err = sql.Open("mysql", connection)

	if err != nil {
		panic(err)
	}

	configureDB(this.db);
}

func (this *DatabaseManager) getItems () *[]Item {
	rows, err := this.db.Query("SELECT * FROM item")

	if (err) != nil {
		panic(err)
	}

	item := Item{}
	result := []Item{}

	for rows.Next() {
		err = rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Price,
		)

		if err != nil {
			panic(err)
		}

		result = append(result, item)
	}

	return &result
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
