package usecase

import "excelToDb/domain"

// DataUsecase представляет use case для работы с данными.
type DataUsecase struct {
	repo domain.DataRepository
}

// NewDataUsecase создает новый экземпляр DataUsecase.
func NewDataUsecase(repo domain.DataRepository) *DataUsecase {
	return &DataUsecase{repo: repo}
}

// SaveDataToDB сохраняет данные в базе данных.
func (uc *DataUsecase) SaveDataToDB(data []domain.Data) error {
	return uc.repo.SaveData(data)
}

// GetDataFromDB получает данные из базы данных.
func (uc *DataUsecase) GetDataFromDB() ([]domain.Data, error) {
	return uc.repo.GetData()
}
