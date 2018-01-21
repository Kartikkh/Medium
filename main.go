package main

import (
	"net/http"
	"os"
	"github.com/kartikkh/Medium/handlers"
	"github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/kartikkh/Medium/config"
    "github.com/mattn/go-colorable"
    "github.com/kartikkh/Medium/models"
)


const (
	DATABASE string = "medium"
	DIALECT  string = "mysql"
)



func main() {

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())

	logrus.Info("succeeded")
	logrus.Warn("not correct")

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	db, err := models.NewDB(DIALECT, DATABASE)

	if err != nil {
		logrus.Fatal(err)
	}
	db.InitSchema()

	h := handlers.NewHandler(db)

    http.HandleFunc("/api/user" ,h.UsersHandler )
	http.HandleFunc("/api/users/login", h.LoginHandler)
	http.HandleFunc("/api/users", h.RegisterHandler)
	http.HandleFunc("/api/profiles/",h.ProfileHandler)
	//http.HandleFunc("/api/articles",h.ProfileHandler)

	logrus.Info("Starting Server on http://localhost:", os.Getenv("port"))

	error := http.ListenAndServe(config.LoadConfig().Host+ ":" + config.LoadConfig().Port, nil)

	if error != nil {
		logrus.Fatal("Error Connecting to Server ! ")
	}

}
