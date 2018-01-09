package main

import (
	"net/http"
	"fmt"
	"os"
	"github.com/julienschmidt/httprouter"
	//"github.com/kartikkh/Medium/controllers"
	//"github.com/kartikkh/Medium/models"
	"github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"github.com/kartikkh/Medium/config"
    "github.com/mattn/go-colorable"
)

func Init(){

	dbuser := os.Getenv("dbUser")
	dbpassword := os.Getenv("dbPassword")
	dbAddr := "mongodb://" + dbuser + ":" + dbpassword + "@ds163745.mlab.com:63745/medium"

	_, err := mgo.Dial(dbAddr)
	if err != nil {
		logrus.Fatal("Error connecting DataBase")
	}
}

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())

	logrus.Info("succeeded")
	logrus.Warn("not correct")

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	Init()
	r := httprouter.New()

	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome kartik!\n")
	})



	logrus.Info("Starting Server on http://localhost:", os.Getenv("port"))
	error := http.ListenAndServe(config.LoadConfig().Host+ ":" + config.LoadConfig().Port, r)
	if error != nil {
		logrus.Fatal("Error Connecting to Server ! ")
	}




}
