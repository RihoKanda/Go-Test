package main

import (
	"log"

	//"github.com/saf1o/go-test/internal/controller"

	"github.com/saf1o/go-test/internal/model"
)

func main() {
	dsn := ""
	if err := model.InitDB(dsn); err != nil {
		log.Fatal(err)
	}

	log.Panicln("DB connected")

	//controller.InitRouter()
	//log.Panicln("server start :8080")
	//http.ListenAndServe(":8080", nil)
}
