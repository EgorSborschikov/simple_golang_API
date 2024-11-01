package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBooks)

	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверка, что ответ является JSON
	if !json.Valid(rr.Body.Bytes()) {
		t.Errorf("handler returned non-JSON body: %s", rr.Body.String())
	}
}

func TestCreateBook(t *testing.T) {
	book := Book{
		Title: "Test Book",
		Author: &Author{
			Firstname: "Test",
			Lastname:  "Author",
		},
	}

	bookJSON, err := json.Marshal(book)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(bookJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createBook)

	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверка, что ответ является JSON
	if !json.Valid(rr.Body.Bytes()) {
		t.Errorf("handler returned non-JSON body: %s", rr.Body.String())
	}
}

func TestGetBook(t *testing.T) {
	// Сначала создаем книгу
	book := Book{
		Title: "Test Book",
		Author: &Author{
			Firstname: "Test",
			Lastname:  "Author",
		},
	}

	books = append(books, book) // Добавляем книгу в глобальный список

	req, err := http.NewRequest("GET", "/books/"+book.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBook)

	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверка, что ответ является JSON
	if !json.Valid(rr.Body.Bytes()) {
		t.Errorf("handler returned non-JSON body: %s", rr.Body.String())
	}
}

func TestUpdateBook(t *testing.T) {
	// Сначала создаем книгу
	book := Book{
		Title: "Test Book",
		Author: &Author{
			Firstname: "Test",
			Lastname:  "Author",
		},
	}

	books = append(books, book) // Добавляем книгу в глобальный список

	updatedBook := Book{
		Title: "Updated Test Book",
		Author: &Author{
			Firstname: "Updated",
			Lastname:  "Author",
		},
	}

	bookJSON, err := json.Marshal(updatedBook)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/books/"+book.ID, bytes.NewBuffer(bookJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateBook)

	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверка, что ответ является JSON
	if !json.Valid(rr.Body.Bytes()) {
		t.Errorf("handler returned non-JSON body: %s", rr.Body.String())
	}
}

func TestDeleteBook(t *testing.T) {
	// Сначала создаем книгу
	book := Book{
		ID: "1",
		Title: "Test Book",
		Author: &Author{
			Firstname: "Test",
			Lastname:  "Author",
		},
	}

	// Добавляем книгу в глобальный список
	books = append(books, book)

	// Создаем запрос на удаление книги
	req, err := http.NewRequest("DELETE", "/books/"+book.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteBook)

	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверка, что книга была удалена
	for _, item := range books {
		if item.ID == book.ID {
			t.Errorf("book with ID %v was not deleted", book.ID)
		}
	}
}