package xls

import (
	"fmt"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"github.com/xuri/excelize/v2"
)

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

var regionalSalutation = map[string]string{
	"Baden-Württemberg":      "Grüßle",
	"Bayern":                 "Servus",
	"Berlin":                 "Juten Tach",
	"Brandenburg":            "Hallöchen",
	"Bremen":                 "Moin",
	"Hamburg":                "Moin Moin",
	"Hessen":                 "Gude",
	"Mecklenburg-Vorpommern": "Moin",
	"Niedersachsen":          "Moin",
	"Nordrhein-Westfalen":    "Tach auch",
	"Rheinland-Pfalz":        "Hallo",
	"Saarland":               "Guddn",
	"Sachsen":                "Glück auf",
	"Sachsen-Anhalt":         "Hallöchen",
	"Schleswig-Holstein":     "Moin",
	"Thüringen":              "Guten Tach",
}

func (g *Generator) Generate(outputPath string, politicians []v1.Politician) error {
	// Create a new Excel file
	f := excelize.NewFile()

	// Define the headers
	headers := []string{"Firma", "Firma Zusatz", "Anrede", "Titel", "Vorname", "Nachname", "Strasse", "Hausnummer", "Postfach", "PLZ", "Ort", "Land", "Briefanrede"}

	// Add the headers to the first row
	sheetName := "Sheet1"
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i))) // Map A, B, C, etc.
		if err := f.SetCellValue(sheetName, cell, header); err != nil {
			return fmt.Errorf("failed to set header %q: %w", header, err)
		}
	}

	// Populate data rows
	for rowIndex, politician := range politicians {

		salutation := fmt.Sprintf("%s, wir hätten da mal ein paar Ideen: https://www.buntesdach.de/p/%s", regionalSalutation[politician.Bio.State], politician.Bio.Id.Value)

		if len(salutation) > 80 {
			return fmt.Errorf("salutation too long: %s", salutation)
		}

		var gender string
		if politician.Bio.Gender == "Männlich" {
			gender = "Herr"
		} else if politician.Bio.Gender == "Weiblich" {
			gender = "Frau"
		}

		rowNum := rowIndex + 2 // Data starts from row 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowNum), "Deutscher Bundestag")
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowNum), "")
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", rowNum), gender)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", rowNum), politician.Bio.AcademicTitle)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", rowNum), politician.Bio.FirstName)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", rowNum), politician.Bio.LastName)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", rowNum), "Platz der Republik")
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", rowNum), "1")
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", rowNum), "")
		f.SetCellValue(sheetName, fmt.Sprintf("J%d", rowNum), "11011")
		f.SetCellValue(sheetName, fmt.Sprintf("K%d", rowNum), "Berlin")
		f.SetCellValue(sheetName, fmt.Sprintf("L%d", rowNum), "Deutschland")
		f.SetCellValue(sheetName, fmt.Sprintf("M%d", rowNum), salutation)
	}

	// Save the file
	if err := f.SaveAs(outputPath); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}
