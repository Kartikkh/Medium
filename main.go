package main

import (
	"net/http"
	"log"
	"fmt"
	"os"
	"github.com/julienschmidt/httprouter"
	//"github.com/kartikkh/Medium/controllers"
	//"github.com/kartikkh/Medium/models"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"github.com/kartikkh/Medium/config"
)


func Init(){

	dbuser := os.Getenv("dbUser")
	dbpassword := os.Getenv("dbPassword")
	dbAddr := "mongodb://" + dbuser + ":" + dbpassword + "@ds163745.mlab.com:63745/medium"

	session, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Fatal("Error connecting DataBase")
	}
}


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	//logger := log.New(os.Stdout,"",1)

	Init()

	r := httprouter.New()

	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})




	error := http.ListenAndServe(config.LoadConfig().Host+ ":" + config.LoadConfig().Port, r)
	if error != nil {
		log.Fatal("Error Connectiong to Server ! ")
	}


}
