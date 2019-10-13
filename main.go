package main

func main () {
	database := new(DatabaseManager)
	database.Start()

	webserver := new(Webserver)
	webserver.Start()
}

