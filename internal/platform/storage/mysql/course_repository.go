package mysql

import (
	mooc "codelytv-api/internal"
	"context"
	"database/sql"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	query := "INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		course.ID(),
		course.Name(),
		course.Duration(),
	)

	return err
}
