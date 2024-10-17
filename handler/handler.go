package handler

import (
	"SAU-TEST/config"
	"SAU-TEST/model"
	"SAU-TEST/payload"
	"SAU-TEST/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []model.Category
	db := config.ConnectDB()

	query := "select * from category"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch category", http.StatusNotFound)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.Id, &category.Name); err != nil {
			http.Error(w, "Failed to scan category", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	utils.WriteToResponseBody(w, categories, http.StatusOK)
}

func PostCategories(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	if r.Method != http.MethodPost {
		utils.WriteToResponseBody(w, map[string]string{"error": "Invalid request method"}, http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&category)
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Failed to parse request body"}, http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	query := "INSERT INTO category (name) VALUES (?)"
	result, err := db.Exec(query, category.Name)
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Failed to create category"}, http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Failed to retrieve category ID"}, http.StatusInternalServerError)
		return
	}

	category.Id = int(id)
	utils.WriteToResponseBody(w, category, http.StatusCreated)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	var items []payload.ItemPayload
	db := config.ConnectDB()

	query := `SELECT 
            item.id, 
            item.name, 
            item.category_id, 
            item.description, 
            item.created_at, 
            item.price, 
            category.id AS category_id, 
            category.name AS category_name
        FROM 
            item 
        JOIN 
            category ON item.category_id = category.id`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch item", http.StatusNotFound)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var item payload.ItemPayload
		if err := rows.Scan(&item.Id,
			&item.Name,
			&item.CategoryId,
			&item.Description,
			&item.CreatedAt,
			&item.Price,
			&item.CategoryName,
		); err != nil {
			http.Error(w, "Failed to scan category", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	utils.WriteToResponseBody(w, items, http.StatusOK)

}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr) // Convert the string ID to an integer
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Invalid item ID"}, http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	// SQL query to get the item by ID
	query := `
        SELECT 
            item.id, 
            item.name, 
            item.category_id, 
            item.description, 
            item.created_at, 
            item.price, 
            category.id AS category_id, 
            category.name AS category_name
        FROM 
            item 
        JOIN 
            category ON item.category_id = category.id
        WHERE 
            item.id = ?`
	var item payload.ItemPayload
	row := db.QueryRow(query, id)
	err = row.Scan(
		&item.Id,
		&item.Name,
		&item.CategoryId,
		&item.Description,
		&item.CreatedAt,
		&item.Price,
		&item.CategoryName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteToResponseBody(w, map[string]string{"error": "Item not found"}, http.StatusNotFound)
		} else {
			utils.WriteToResponseBody(w, map[string]string{"error": "Failed to fetch item"}, http.StatusInternalServerError)
		}
		return
	}
	utils.WriteToResponseBody(w, item, http.StatusOK)
}

func PostItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteToResponseBody(w, map[string]string{"error": "Invalid request method"}, http.StatusMethodNotAllowed)
		return
	}

	var itemPayload payload.ItemPost
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&itemPayload)
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Failed to parse request body"}, http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	query := "INSERT INTO item (category_id, name, description, price) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, itemPayload.CategoryId, itemPayload.Name, itemPayload.Description, itemPayload.Price)
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Failed to create item"}, http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Failed to retrieve item ID"}, http.StatusInternalServerError)
		return
	}

	item := payload.ItemPost{
		Id:          int(id),
		CategoryId:  itemPayload.CategoryId,
		Name:        itemPayload.Name,
		Description: itemPayload.Description,
		Price:       itemPayload.Price,
	}

	utils.WriteToResponseBody(w, item, http.StatusCreated)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr) // Convert the string ID to an integer
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Invalid item ID"}, http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()
	query := "delete from item where id =? "
	_, errs := db.Exec(query, id)
	if errs != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Failed to delete item"}, http.StatusInternalServerError)
		return
	}
	utils.WriteToResponseBody(w, map[string]string{"message": "Item deleted successfully"}, http.StatusOK)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr) 
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Invalid item ID"}, http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodPut {
		utils.WriteToResponseBody(w, map[string]string{"error": "Invalid request method"}, http.StatusMethodNotAllowed)
		return
	}

	var itemPayload payload.ItemPut
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&itemPayload)
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Failed to parse request body"}, http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	query := "UPDATE item SET name = ?, description = ?, price = ? WHERE id = ?"
	_, err = db.Exec(query, itemPayload.Name, itemPayload.Description, itemPayload.Price, id)
	if err != nil {
		utils.WriteToResponseBody(w, map[string]string{"error": "Failed to update item"}, http.StatusInternalServerError)
		return
	}

	utils.WriteToResponseBody(w, map[string]string{"message": "Item updated successfully"}, http.StatusOK)
}
