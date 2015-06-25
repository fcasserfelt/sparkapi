package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/thoas/stats"
)

// Http server port
const Port = "5050"

func main() {
	statsMiddleware := stats.New()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})

	n := negroni.Classic()

	router := mux.NewRouter()
	apiRoutes := ApiRouter(statsMiddleware)

	router.HandleFunc("/", HomeHandler)
	router.PathPrefix("/api").Handler(negroni.New(
		statsMiddleware,
		negroni.Wrap(apiRoutes),
	))

	n.Use(c)
	n.UseHandler(router)
	n.Run("0.0.0.0:" + Port)

}

// HomeHandler A simple message on root path
func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Spark API")
}

// ToJSON Converts object to JSON string
func ToJSON(obj interface{}) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("[ERR] ToJSON - ", err)
		return "{}", err
	}
	return string(bytes), nil
}
