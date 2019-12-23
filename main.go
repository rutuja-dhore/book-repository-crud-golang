package main
import (
	"log"
	"net/http"
	"fmt"
	"rsc.io/quote/v3"
	"golangproject/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
	fmt.Println(quote.HelloV3())

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8084", r))
}