// main.go
package main

// @title ExcelToDB API
// @version 1.0
// @description This is an API for reading data from an Excel file, saving it to a database, and fetching data from the database.
// @host localhost:8080
// @BasePath /

import (
	"database/sql"
	"excelToDb/internal/domain"
	"excelToDb/internal/handler"
	"github.com/gin-gonic/gin"
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

	// Создаем новый экземпляр DataUsecase
	dataUsecase := domain.NewDataUsecase(domain.NewMySQLRepository())

	// Создаем новый экземпляр DataHandler
	dataHandler := handler.NewDataHandler(*dataUsecase)

	r := gin.Default()

	// Маршрутизируем GET-запрос на обработчик чтения данных из Excel-файла
	r.GET("/read", dataHandler.ReadData)

	// Маршрутизируем POST-запрос на обработчик записи данных в базу данных
	r.POST("/write", dataHandler.WriteData)

	r.Run(":8080")
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
