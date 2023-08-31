package postgres

import (
	"fmt"
)

func CreateUser(s *Storage, user_name string) (int64, error) {
	const op = "postgres.users.CreateUser"

	var id int64
	err := s.db.QueryRow("INSERT INTO users(name) VALUES ($1) RETURNING id;", user_name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func DeleteUser(s *Storage, user_name string) (int64, error) {
	const op = "postgres.segments.DeleteSegment"
	
	stmt, err := s.db.Prepare("DELETE FROM users WHERE name = ($1);")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	_ , err = stmt.Exec(user_name)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return 0, nil
}
