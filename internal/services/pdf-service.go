package services

import (
	"bytes"
	"fmt"
	"time"

	"github.com/ak-ansari/go-pdf-service/pkg/model"
	"github.com/jung-kurt/gofpdf"
)

func GenerateStudentReport(student *model.Student) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Header Of PDF
	pdf.SetFont("Arial", "B", 18)
	pdf.Cell(0, 10, "Student Report")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 8, "Academic Management System")
	pdf.Ln(15)

	section := func(title string) {
		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(0, 10, title)
		pdf.Ln(10)
		pdf.SetFont("Arial", "", 11)
	}

	row := func(label, value string) {
		pdf.CellFormat(50, 8, label, "", 0, "", false, 0, "")
		pdf.CellFormat(0, 8, value, "", 1, "", false, 0, "")
	}

	// Student Info
	section("Student Information")
	row("Name", student.Name)
	row("Class", value(student.Class))
	row("Section", value(student.Section))
	row("Roll No", intValue(student.Roll))

	// Personal Details
	section("Personal Details")
	row("Email", student.Email)
	row("Phone", value(student.Phone))
	row("Gender", value(student.Gender))
	row("Date of Birth", dateValue(student.DOB))

	// Family Details
	section("Family Details")
	row("Father Name", value(student.FatherName))
	row("Father Phone", value(student.FatherPhone))
	row("Mother Name", value(student.MotherName))
	row("Mother Phone", value(student.MotherPhone))
	row("Guardian Name", value(student.GuardianName))
	row("Guardian Phone", value(student.GuardianPhone))
	row("Relation", value(student.RelationOfGuardian))

	// Address
	section("Address")
	row("Current Address", value(student.CurrentAddress))
	row("Permanent Address", value(student.PermanentAddress))

	// Administrative
	section("Administrative Details")
	row("Admission Date", dateValue(student.AdmissionDate))
	row("Class Teacher", value(student.ReporterName))
	row("System Access", boolValue(student.SystemAccess))

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func value(v *string) string {
	if v == nil || *v == "" {
		return "N/A"
	}
	return *v
}

func intValue(v *int) string {
	if v == nil {
		return "N/A"
	}
	return fmt.Sprintf("%d", *v)
}

func dateValue(v *time.Time) string {
	if v == nil {
		return "N/A"
	}
	return v.Format("02 Jan 2006")
}

func boolValue(v bool) string {
	if v {
		return "Active"
	}
	return "Inactive"
}
