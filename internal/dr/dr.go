package dr

import "database/sql"

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
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	d.db = db
	return nil
}

func (d *DR) Close() {

}
