package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ak-ansari/go-pdf-service/internal/client"
	"github.com/ak-ansari/go-pdf-service/internal/handler"
	"github.com/ak-ansari/go-pdf-service/internal/route"
	"github.com/gorilla/mux"
)

func main() {
	baseURL := os.Getenv("STUDENT_API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:5007/api/v1"
	}

	studentClient := client.NewStudentAPIClient(baseURL)
	reportHandler := handler.NewReportHandler(studentClient)

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	route.RegisterReportRoute(r,reportHandler)

	log.Println("Go PDF service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
