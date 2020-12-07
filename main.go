package main

import (
	"context"
	"net/http"
	"os"

	"github.com/ademcaglin/authserver/handlers"
	"github.com/ademcaglin/authserver/models"
	"github.com/ademcaglin/authserver/stores/mongodb"
)

func main() {
	client := mongodb.GetClient(os.Getenv("MONGO_CONSTR"))
	userStore := mongodb.NewMongoUserStore(client)
	store := models.Store{Users: userStore}
	defer client.Disconnect(context.Background())
	fs := http.FileServer(http.Dir("./templates/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/authorize", handlers.AuthorizeHandler(&store))
	http.ListenAndServe(":8000", nil)
}
