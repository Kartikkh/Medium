package main

import (
	"net/http"
	"os"
	//"github.com/kartikkh/Medium/controllers"
	"github.com/kartikkh/Medium/handlers"
	"github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"github.com/kartikkh/Medium/config"
    "github.com/mattn/go-colorable"
)

func getDatabase() *mgo.Session{

	dbuser := os.Getenv("dbUser")
	dbpassword := os.Getenv("dbPassword")
	dbAddr := "mongodb://" + dbuser + ":" + dbpassword + "@ds163745.mlab.com:63745/medium"

	session , err := mgo.Dial(dbAddr)

	if err != nil {
		logrus.Fatal("Error connecting DataBase")
	}
	return session

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



	h := handlers.NewHandler(getDatabase())

    http.HandleFunc("/api/user" ,h.UsersHandler )
	http.HandleFunc("/api/users/login", h.LoginHandler)
	http.HandleFunc("/api/users", h.RegisterHandler)

	logrus.Info("Starting Server on http://localhost:", os.Getenv("port"))

	error := http.ListenAndServe(config.LoadConfig().Host+ ":" + config.LoadConfig().Port, nil)

	if error != nil {
		logrus.Fatal("Error Connecting to Server ! ")
	}




}
