package main

import (
	"net/http"

	"github.com/ademcaglin/authserver/handlers"
)

/*func printUser(store models.UserStore) {
	x, _ := store.GetOne(context.TODO(), "ademcaglin")
	fmt.Println(x)
}
func main() {
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoServer.Stop()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoServer.URI()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	store := mongodb.NewMongoUserStore(client)
	store.Save(context.TODO(), "ademcaglin", "Adem Çağlın")
	printUser(store)
}*/

func main() {
	fs := http.FileServer(http.Dir("./templates/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handlers.HomeHandler)
	//http.HandleFunc("/authorize", handlers.AuthorizeHandler(store))
	http.ListenAndServe(":8000", nil)
}
