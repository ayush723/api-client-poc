package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// swagger:model Book
type Book struct {
	// ID is id of book
	// required: true
	ID string `json:"id,omitempty"`
	// Title is title of book
	// required: true
	Title string `json:"title,omitempty"`
	// Author is author of the book
	// required: true
	Author string `json:"author,omitempty"`
	// Price is the price of book
	// required: true
	Price float64 `json:"price,omitempty"`
}

var books []Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /books GetBooks
	//
	// Returns a list of books
	//
	//   Produces:
	//   application/json
	//
	//   Responses:
	//     200: GetBooksResponse

	// swagger:response GetBooksResponse
	type Response struct {
		// in: body
		// the books for this user
		Books []Book `json:"books"`
	}

	var boos []Book
	boos = append(boos, Book{ID: "1", Title: "Book 1", Author: "Author 1", Price: 9.99})
	boos = append(boos, Book{ID: "2", Title: "Book 2", Author: "Author 2", Price: 14.99})
	fmt.Printf("returned:%v\n", boos)
	json.NewEncoder(w).Encode(Response{Books: boos})
}

func main() {
	router := mux.NewRouter()
	v1 := router.PathPrefix("/v1").Subrouter()
	books = append(books, Book{ID: "1", Title: "Book 1", Author: "Author 1", Price: 9.99})
	books = append(books, Book{ID: "2", Title: "Book 2", Author: "Author 2", Price: 14.99})

	v1.HandleFunc("/books", GetBooks).Methods(http.MethodGet)

	fmt.Println("listening at :9000 ...")
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(credentials, methods, origins)(router)))
}
