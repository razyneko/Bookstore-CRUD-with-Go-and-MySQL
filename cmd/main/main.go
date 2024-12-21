package main
//  here we will create server and tell go where our routers are
import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/razyneko/Bookstore-Management-Go-MySQL/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}