package route

import (
	"github.com/ak-ansari/go-pdf-service/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterReportRoute(r *mux.Router,handler *handler.ReportHandler){
	r.HandleFunc("/students/{id}/report",handler.GetStudentReport).Methods("GET")
}