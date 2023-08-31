package postgres

import (
	"fmt"
)

func (s *Storage) AddSegmentToUser(name string, user_id int) error {

	err := s.db.QueryRow("INSERT INTO users_segments (user_id, segment_id) VALUES ($1, (SELECT id AS segment_id FROM segments WHERE name = $2)) RETURNING id;", user_id, name)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (s *Storage) RemoveSegmentFromUser (name string, user_id int) error {

	err := s.db.QueryRow("DELETE FROM users_segments WHERE user_id = $1 AND segment_id IN (SELECT id FROM segments WHERE name = $2);", user_id, name)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (s *Storage) GetActiveSegments(user_id int) ([]string, error) {

	var ActiveSegments []string

	res, err := s.db.Query("SELECT segments.name FROM users_segments INNER JOIN segments ON users_segments.segment_id = segments.id WHERE user_id = $1;", user_id)
	_ = err
	for res.Next() {
        var segment string
        err = res.Scan(&segment)
        if err != nil {
            fmt.Println("failed to scan", err)
        }
		ActiveSegments = append(ActiveSegments, segment)
	}
	return ActiveSegments, nil
}


	