package config

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB{
	DB, err := sql.Open("mysql", "root:@/SAU-Test")
	if err != nil {
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil
	}

	log.Println("Database connected!")
	return DB
}

func CreateTables(db *sql.DB) {
    createCategoryTable := `
    CREATE TABLE IF NOT EXISTS category (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL UNIQUE
    );`

    createItemTable := `
    CREATE TABLE IF NOT EXISTS item (
        id INT AUTO_INCREMENT PRIMARY KEY,
        category_id INT,
        name VARCHAR(255) NOT NULL,
        description TEXT,
        price DECIMAL(10, 2) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (category_id) REFERENCES category(id)
    );`

    // Execute the queries
    if _, err := db.Exec(createCategoryTable); err != nil {
        log.Fatal("Failed to create category table:", err)
    }

    if _, err := db.Exec(createItemTable); err != nil {
        log.Fatal("Failed to create item table:", err)
    }

    log.Println("Tables created successfully")
}