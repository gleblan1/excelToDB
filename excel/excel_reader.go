package excel

import (
	"github.com/xuri/excelize/v2"

	"excelToDb/domain"
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
		if len(row) >= 2 {
			d := domain.Data{
				Column1: row[0],
				Column2: row[1],
			}
			data = append(data, d)
		}
	}

	return data, nil
}
