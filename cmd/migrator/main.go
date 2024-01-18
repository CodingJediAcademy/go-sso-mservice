package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	// Получаем необходимые значения из флагов запуска

	// Путь до файла БД.
	// Его достаточно, т.к. мы используем SQLite, другие креды не нужны.
	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	// Путь до папки с миграциями.
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	// Таблица, в которой будет храниться информация о миграциях. Она нужна
	// для того, чтобы понимать, какие миграции уже применены, а какие нет.
	// Дефолтное значение - 'migrations'.
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse() // Выполняем парсинг флагов

	// Валидация параметров
	if storagePath == "" {
		// Простейший способ обработки ошибки :)
		// При необходимости, можете выбрать более подходящий вариант.
		// Меня паника пока устраивает, поскольку это вспомогательная утилита.
		panic("storage-path is required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	// Выполняем миграции до последней версии
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}
}
