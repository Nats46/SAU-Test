package handler

import (
    "SAU-TEST/config"
    "SAU-TEST/model"
    "SAU-TEST/payload"
    "bytes"
    "database/sql"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

var db *sql.DB

func setup() {
    db = config.ConnectDB()
}

func teardown() {
}

func TestGetCategories(t *testing.T) {
    setup()
    defer teardown()

    req, err := http.NewRequest("GET", "/categories", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetCategories)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var categories []model.Category
    json.Unmarshal(rr.Body.Bytes(), &categories)

}

func TestPostCategories(t *testing.T) {
    setup()
    defer teardown()

    newCategory := model.Category{Name: "Test Category"}
    payloadBytes, _ := json.Marshal(newCategory)

    req, err := http.NewRequest("POST", "/categories/new", bytes.NewBuffer(payloadBytes))
    if err != nil {
        t.Fatal(err)
    }

    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(PostCategories)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }

}

func TestGetItems(t *testing.T) {
    setup()
    defer teardown()

    req, err := http.NewRequest("GET", "/items", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetItems)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

}

func TestGetItemById(t *testing.T) {
    setup()
    defer teardown()

    req, err := http.NewRequest("GET", "/items/?id=1", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetItemById)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

}

func TestPostItem(t *testing.T) {
    setup()
    defer teardown()

    newItem := payload.ItemPost{
        CategoryId: 1,
        Name:       "New Item",
        Description: "Description for new item",
        Price:      99.99,
    }
    payloadBytes, _ := json.Marshal(newItem)

    req, err := http.NewRequest("POST", "/items/new", bytes.NewBuffer(payloadBytes))
    if err != nil {
        t.Fatal(err)
    }

    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(PostItem)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }

}

func TestDeleteItem(t *testing.T) {
    setup()
    defer teardown()

    req, err := http.NewRequest("DELETE", "/items/delete/?id=1", nil) // Assume item ID 1 exists
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(DeleteItem)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

}

func TestUpdateItem(t *testing.T) {
    setup()
    defer teardown()

    updatedItem := payload.ItemPut{
        Name:        "Updated Item",
        Description: "Updated description",
        Price:       89.99,
    }
    payloadBytes, _ := json.Marshal(updatedItem)

    req, err := http.NewRequest("PUT", "/items/update/?id=1", bytes.NewBuffer(payloadBytes)) // Assume item ID 1 exists
    if err != nil {
        t.Fatal(err)
    }

    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(UpdateItem)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}
