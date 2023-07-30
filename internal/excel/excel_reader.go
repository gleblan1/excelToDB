package excel

import (
	"github.com/xuri/excelize/v2"

	"excelToDb/internal/domain"
)

// ReadFromExcelFile считывает данные из файла Excel и возвращает их в виде среза структур domain.Data и ошибку.
func ReadFromExcelFile(filePath string) ([]domain.Data, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	// Здесь предполагается, что данные хранятся в первом листе (Sheet1) в столбцах A и B.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	var data []domain.Data
	for _, row := range rows {
		var d domain.Data
		if len(row) >= 1 {
			d.Column1 = row[0]
		}
		if len(row) >= 2 {
			d.Column2 = row[1]
		}
		data = append(data, d)
	}

	return data, nil
}
