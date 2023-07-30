package domain

// Data представляет структуру данных, которые будут записаны в Excel-файл и базу данных.
type Data struct {
	Column1 string
	Column2 string
}

// DataRepository определяет интерфейс для работы с данными в базе данных.
type DataRepository interface {
	SaveData(data []Data) error
	GetData() ([]Data, error)
}
