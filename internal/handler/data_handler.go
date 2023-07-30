package handler

import (
	"excelToDb/internal/domain"
	"excelToDb/internal/excel"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataHandler struct {
	dataUsecase   domain.DataUsecase
	dataFromExcel []domain.Data // Поле для хранения данных из ReadData
}

func NewDataHandler(dataUsecase domain.DataUsecase) *DataHandler {
	return &DataHandler{dataUsecase: dataUsecase}
}

func (h *DataHandler) ReadData(c *gin.Context) {
	dataFromExcel, err := excel.ReadFromExcelFile("input.xlsx")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения данных из Excel-файла"})
		return
	}

	h.dataFromExcel = dataFromExcel // Сохраняем данные в поле dataFromExcel

	c.JSON(http.StatusOK, dataFromExcel)
}

// WriteData сохраняет данные в базе данных и создает Excel-файл.
func (h *DataHandler) WriteData(c *gin.Context) {
	if h.dataFromExcel == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Отсутствуют данные для записи в базу данных"})
		return
	}

	// Используем данные из h.dataFromExcel для записи в базу данных
	if err := h.dataUsecase.SaveDataToDB(h.dataFromExcel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка записи данных в базу данных"})
		return
	}

	filePath := "output.xlsx"
	if err := excel.WriteToExcelFile(h.dataFromExcel, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка записи данных в Excel-файл"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Данные успешно записаны в базу данных и создан Excel-файл"})
}
