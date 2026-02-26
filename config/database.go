package config

import "fmt"

type Database struct {
	DBName   string `mapstructure:"dbname"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	SSLMode  bool   `mapstructure:"sslmode"`
}

func (db *Database) DSN() string {
	sslmode := "disable"

	if db.SSLMode {
		sslmode = "enable"
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host,
		db.Port,
		db.Username,
		db.Password,
		db.DBName,
		sslmode,
	)
}

func (db *Database) IsValid() bool {
	return db.Username != "" &&
		db.Password != "" &&
		db.Host != "" &&
		db.Port != "" &&
		db.DBName != ""
}
