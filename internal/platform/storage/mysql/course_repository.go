package mysql

import (
	mooc "codelytv-api/internal/mooc"
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
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
	sqlCourseStruct := sqlbuilder.NewStruct(new(sqlCourse))

	var (
		id       = course.ID()
		name     = course.Name()
		duration = course.Duration()
	)

	query, args := sqlCourseStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       id.Value(),
		Name:     name.Value(),
		Duration: duration.Value(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course: %v", err)
	}

	return nil
}
