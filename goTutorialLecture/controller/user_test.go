package controller

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	res, err := http.Get("http://localhost:8080/users")
	fmt.Println("user testing", res.StatusCode)
	if http.StatusOK != res.StatusCode {
		t.Errorf("Expected response code %d. got %d\n", http.StatusOK, res.StatusCode)
	}
	if err != nil {
		t.Errorf("encountered an error %s", err)
	}
}

func TestGetProfile(t *testing.T) {
	res, err := http.Get("http://localhost:8080/profile")
	fmt.Println("profile testing", res.StatusCode)
	if http.StatusOK != res.StatusCode {
		t.Errorf("Expected response code %d. got %d\n", http.StatusOK, res.StatusCode)
	}
	if err != nil {
		t.Errorf("encountered an error %s", err)
	}
}
