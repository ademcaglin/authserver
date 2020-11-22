package handlers

import (
	"net/http"

	"github.com/ademcaglin/authserver/models"
)

func AuthorizeHandler(store *models.Store) func(http.ResponseWriter, *http.Request) {
	if store == nil {
		panic("nil Store!")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			get(store)
		}
	}
}

func get(store *models.Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Users.GetOne(nil, "")
	}
}
