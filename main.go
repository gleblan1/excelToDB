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
		log.Fatal("Ошибка чтения конфига", err.Error())
	}

	// Устанавливаем соединение с базой данных
	db, err := sql.Open("mysql", viper.GetString("dbOpen"))

	if err != nil {
		log.Fatal("Ошибка открытия базы данных:", err)
	}
	defer db.Close()

	// Откатываем миграции вниз перед применением
	err = rollbackMigrations(db)
	if err != nil {
		log.Fatal("Ошибка отката миграций:", err)
	}

	// Применяем миграции вверх
	err = applyMigrations(db)
	if err != nil {
		log.Fatal("Ошибка применения миграций:", err)
	}

	// Чтение данных из Excel-файла
	dataFromExcel, err := excel.ReadFromExcelFile("input.xlsx")
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
	err = excel.WriteToExcelFile(dataFromDB, "output.xlsx")
	if err != nil {
		log.Fatal("Ошибка записи данных в Excel-файл:", err)
	}

	log.Println("Программа успешно завершена.")
}

// applyMigrations применяет миграции базы данных вверх
func applyMigrations(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	_, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	return err
}

// rollbackMigrations откатывает все миграции базы данных вниз
func rollbackMigrations(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	_, err := migrate.ExecMax(db, "mysql", migrations, migrate.Down, 0)
	return err
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
