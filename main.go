package main

import (
	"SAU-TEST/config"
	"SAU-TEST/route"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 1 {
        if os.Args[1] == "migrate" {
			db:=config.ConnectDB()
            config.CreateTables(db)
            log.Println("Migration completed!")
            return
        }
    }else{
		route.SetupRoutes() 
		log.Println("Starting the server on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))	
	}
}