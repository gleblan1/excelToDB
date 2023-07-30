package excel

import (
	"excelToDb/internal/domain"
	"github.com/xuri/excelize/v2"
	"strconv"
)

// WriteToExcelFile записывает данные в файл Excel.
func WriteToExcelFile(data []domain.Data, filePath string) error {
	f := excelize.NewFile()

	// Создаем новый лист в файле Excel.
	sheetName := "Sheet1"
	index, _ := f.NewSheet(sheetName)

	// Записываем данные в столбцы A и B на листе.
	for i, d := range data {
		rowIndex := i + 1
		f.SetCellValue(sheetName, "A"+strconv.Itoa(rowIndex), d.Column1)
		f.SetCellValue(sheetName, "B"+strconv.Itoa(rowIndex), d.Column2)
	}

	// Устанавливаем активный лист на созданный лист.
	f.SetActiveSheet(index)

	// Сохраняем данные в файл.
	if err := f.SaveAs(filePath); err != nil {
		return err
	}

	return nil
}
