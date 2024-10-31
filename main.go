package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//открытая библотека
	resp, err := http.Get("https://openlibrary.org/subjects/literature.json?limit=5") // Adjust the endpoint as needed
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var data struct {
		Works []struct {
			Title   string `json:"title"`
			Authors []struct {
				Name string `json:"name"`
			} `json:"authors"`
		} `json:"works"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// фрагмент книги
	books = []Book{}
	for _, work := range data.Works {
		author := &Author{Firstname: "", Lastname: ""}
		if len(work.Authors) > 0 {
			author.Firstname = work.Authors[0].Name // Get the first author
		}
		book := Book{
			ID:     strconv.Itoa(rand.Intn(1000000)),
			Title:  work.Title,
			Author: author,
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//чтение из консоли
	reader := bufio.NewReader(os.Stdin)

	// заголовок
	log.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	title = title[:len(title)-1] // Remove the newline character

	// фамилия автора
	log.Print("Enter author's first name: ")
	firstname, _ := reader.ReadString('\n')
	firstname = firstname[:len(firstname)-1] // Удалить символ новой строки

	// имя автора
	log.Print("Enter author's last name: ")
	lastname, _ := reader.ReadString('\n')
	lastname = lastname[:len(lastname)-1] 

	// создать новую книгу
	book := Book{
		ID:    strconv.Itoa(rand.Intn(1000000)),
		Title: title,
		Author: &Author{
			Firstname: firstname,
			Lastname:  lastname,
		},
	}

	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	//инициализация фрагмент книг с некоторыми данными по умолчанию.
	books = append(books, Book{
		ID:     "1",
		Title:  "1984",
		Author: &Author{Firstname: "George", Lastname: "Orwell"},
	})
	books = append(books, Book{
		ID:     "2",
		Title:  "Brave New World",
		Author: &Author{Firstname: "Aldous", Lastname: "Huxley"},
	})

	// пути для запросов 
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// запуск сервера 
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}