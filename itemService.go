package main

type Item struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ItemService struct {
	db *DatabaseManager
}

func (this *ItemService) getItems() *[]Item {
	items := this.db.getItems()
	return items
}
