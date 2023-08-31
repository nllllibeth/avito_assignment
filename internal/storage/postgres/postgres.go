package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg Config) (*Storage, error) {
	connInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to db!")
	return &Storage{db: db}, nil
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}

func CreateTables(s *Storage) {
	_, err := s.db.Query("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL);")
	if err != nil {
		fmt.Println("failed to execute query", err)
		return
	}
	fmt.Println("Table users created successfully")

	_, err = s.db.Query("CREATE TABLE IF NOT EXISTS segments (id SERIAL PRIMARY KEY, name VARCHAR(255) UNIQUE NOT NULL);")
	if err != nil {
		fmt.Println("failed to execute query", err)
		return
	}
	fmt.Println("Table segments created successfully")

	_, err = s.db.Query("CREATE TABLE IF NOT EXISTS users_segments (id SERIAL PRIMARY KEY, user_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL, segment_id INT REFERENCES segments (id) ON DELETE CASCADE NOT NULL, UNIQUE(user_id, segment_id));")
	if err != nil {
		fmt.Println("failed to execute query", err)
		return
	}
	fmt.Println("Table user_segments created successfully")
}
