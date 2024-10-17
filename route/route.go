package route

import (
	"SAU-TEST/handler"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/categories", handler.GetCategories)      // GET /categories
	http.HandleFunc("/categories/new", handler.PostCategories) // POST /categories/new

	http.HandleFunc("/items", handler.GetItems)           // GET /items
	http.HandleFunc("/items/new", handler.PostItem)       // POST /items/new
	http.HandleFunc("/items/", handler.GetItemById)       // GET /items/{id}
	http.HandleFunc("/items/update/", handler.UpdateItem) // PUT /items/update/{id}
	http.HandleFunc("/items/delete/", handler.DeleteItem) // DELETE /items/delete/{id}
}