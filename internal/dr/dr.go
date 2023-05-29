package dr

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DR struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *DR {
	return &DR{
		config: config,
	}
}

func (d *DR) Open() error {
	db, err := sql.Open("postgres", d.config.DatabaseURL)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()

	// Вызов хранимой функции
	rows, err := db.Query("SELECT * FROM public.get_all_orders()")
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса: ", err, db)
		return err
	}
	defer rows.Close()

	// Обработка результатов11
	for rows.Next() {
		var orderUID string
		var trackNumber string
		var dateCreated string
		err = rows.Scan(&orderUID, &trackNumber, &dateCreated)
		if err != nil {
			fmt.Println("Ошибка при сканировании результата:", err)
			return err
		}

		fmt.Printf("Order ID: %s, Order Date: %s, Total Amount: %s\n", orderUID, trackNumber, dateCreated)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Ошибка при получении результатов:", err)
		return err
	}

	d.db = db
	return nil
}

func (d *DR) Close() {

}
