package main

import (
	"assignment1/controller"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetAllUsersController(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/getUsers", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllUsersController)

	//  handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	//here check the response body is not empty
	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "")
	}
}
func TestUpdateUsersController(t *testing.T) {
	requestBody := strings.NewReader(`
		{
			"name":"Test User update",
			"dob":"12-Jan-1998",
			"address":"Ambala",
			"description":"THis is description"
		}
	`) // Create a request to pass to our handler.
	path := fmt.Sprintf("/updateUser/%s", "61d53c12b3c36968d7c78641")
	req, err := http.NewRequest("PUT", path, requestBody)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc("/updateUser/{id}", controller.UpdateUserController)
	// directly and pass in our Request and ResponseRecorder.
	router.ServeHTTP(rr, req)
	
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	//here check the response body is not empty
	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "")
	}
}

func TestInsertUsersController(t *testing.T) {
	requestBody := strings.NewReader(`
		{
			"name":"Test User",
			"dob":"12-Jan-1998",
			"address":"Ambala",
			"description":"THis is description"
		}
	`) // Create a request to pass to our handler.

	req, err := http.NewRequest("POST", "/createUser", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.InsertUserController)

	//  handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	//here check the response body is not empty
	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "")
	}
}

func TestDeleteUserController(t *testing.T) {
	 // Create a request to pass to our handler.
	req, err := http.NewRequest("DELETE", "/deleteUser/61d5a37670145a048ecdecd2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc("/deleteUser/{id}", controller.DeleteUserController)
	// directly and pass in our Request and ResponseRecorder.
	router.ServeHTTP(rr, req)
	
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	//here check the response body is not empty
	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "")
	}
}
