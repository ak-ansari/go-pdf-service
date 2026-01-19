package handler

import (
	"net/http"

	"github.com/ak-ansari/go-pdf-service/internal/client"
	"github.com/ak-ansari/go-pdf-service/internal/services"
	"github.com/gorilla/mux"
)

type ReportHandler struct {
	StudentClient *client.StudentAPIClient
}

func NewReportHandler(c *client.StudentAPIClient) *ReportHandler {
	return &ReportHandler{StudentClient: c}
}

func (h *ReportHandler) GetStudentReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	student, err := h.StudentClient.GetStudentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	pdfBytes, err := services.GenerateStudentReport(student)
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=report_"+student.Name+".pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}
