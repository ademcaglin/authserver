package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ademcaglin/authserver/models"
)

func AuthorizeHandler(store *models.Store) func(http.ResponseWriter, *http.Request) {
	if store == nil {
		panic("nil Store!")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			get(store, w, r)
		}
	}
}

func get(store *models.Store, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, err := store.Users.GetOne(context.Background(), "ademcaglin")
	if err == nil {
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode("adem")
	}
}
