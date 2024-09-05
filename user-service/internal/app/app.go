package app

import (
	"fmt"
	"user-service/internal/database"
)



func Run() {
    fmt.Println("App is running")
    database.Connect()

}