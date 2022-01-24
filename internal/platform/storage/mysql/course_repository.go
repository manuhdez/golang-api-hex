package mysql

import (
	"codelytv-api/internal/mooc"
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

var courseStruct = sqlbuilder.NewStruct(new(sqlCourse))

func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	var (
		id       = course.ID()
		name     = course.Name()
		duration = course.Duration()
	)

	query, args := courseStruct.InsertInto(sqlCourseTable, sqlCourse{
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

func (r *CourseRepository) All(ctx context.Context) ([]mooc.Course, error) {
	qb := sqlbuilder.Select("*").From(sqlCourseTable)
	query, args := qb.Build()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return []mooc.Course{}, fmt.Errorf("could not get the list of courses: %v", err)
	}

	var courses []mooc.Course
	for rows.Next() {
		var row sqlCourse
		err = rows.Scan(courseStruct.Addr(&row)...)
		if err != nil {
			return []mooc.Course{}, fmt.Errorf("error parsing query result: %v", err)
		}

		course, err := mooc.NewCourse(row.ID, row.Name, row.Duration)
		if err == nil {
			courses = append(courses, course)
		}
	}

	return courses, nil
}

func (r *CourseRepository) Find(ctx context.Context, id mooc.CourseID) (mooc.Course, error) {
	qb := sqlbuilder.Select("*").From(sqlCourseTable)
	qb.Where(qb.Equal("id", id.Value()))
	query, args := qb.Build()

	rows, err := r.db.QueryContext(ctx, query, args...)

	if rows.Next() == false {
		return mooc.Course{}, fmt.Errorf("could not find a course with id: %s", id.Value())
	}

	var course sqlCourse
	err = rows.Scan(courseStruct.Addr(&course)...)
	if err != nil {
		return mooc.Course{}, err
	}

	result, err := mooc.NewCourse(course.ID, course.Name, course.Duration)
	if err != nil {
		return mooc.Course{}, err
	}

	return result, nil
}
