package main

import (
	"database/sql"
	"excelToDb/excel"
	datarepository "excelToDb/repository"
	"excelToDb/usecase"
	"github.com/spf13/viper"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rubenv/sql-migrate"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("ошибка в чтении конфиг файла: %s", err.Error())
	}

	// Устанавливаем соединение с базой данных
	db, err := sql.Open("mysql", "root:a910111A!@tcp(127.0.0.1:3306)/excelToDB?parseTime=true")

	if err != nil {
		log.Fatal("Ошибка открытия базы данных:", err)
	}
	defer db.Close()

	// Применяем миграции
	err = applyMigrations(db)
	if err != nil {
		log.Fatal("Ошибка применения миграций:", err)
	}

	// Чтение данных из Excel-файла
	dataFromExcel, err := excel.ReadFromExcelFile("xslx/input.xlsx")
	if err != nil {
		log.Fatal("Ошибка чтения данных из Excel-файла:", err)
	}

	// Создание репозитория и usecase для работы с данными
	repo := datarepository.NewMySQLRepository()
	dataUsecase := usecase.NewDataUsecase(repo)

	// Запись данных в базу данных
	err = dataUsecase.SaveDataToDB(dataFromExcel)
	if err != nil {
		log.Fatal("Ошибка записи данных в базу данных:", err)
	}

	// Получение данных из базы данных
	dataFromDB, err := dataUsecase.GetDataFromDB()
	if err != nil {
		log.Fatal("Ошибка получения данных из базы данных:", err)
	}

	// Запись данных из базы данных в новый Excel-файл
	err = excel.WriteToExcelFile(dataFromDB, "xlsx/output.xlsx")
	if err != nil {
		log.Fatal("Ошибка записи данных в Excel-файл:", err)
	}

	log.Println("Программа успешно завершена.")
}

// applyMigrations применяет миграции базы данных
func applyMigrations(db *sql.DB) error {
	// Инициализируем новый источник миграций из файловой системы
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	// Применяем миграции
	_, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	return err
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
