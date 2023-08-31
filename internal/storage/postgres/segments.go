package postgres

import (
	"fmt"
)
func (s *Storage) CreateSegment(segment_name string) (int64, error) {
	const op = "postgres.segments.CreateSegment"

	var id int64
	err := s.db.QueryRow("INSERT INTO segments(name) VALUES ($1) RETURNING id;", segment_name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *Storage) DeleteSegment(segment_name string) (int64, error) {
	const op = "postgres.segments.DeleteSegment"
	
	stmt, err := s.db.Prepare("DELETE FROM segments WHERE name = ($1);")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	_ , err = stmt.Exec(segment_name)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return 0, nil
}
