package main

func main() {
	database := new(DatabaseManager)
	database.start()

	itemService := &ItemService{db: database}
	controller := &Controller{itemService: itemService}

	webserver := Webserver{controller: controller}
	webserver.start()
}
