package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/manuhdez/golang-api-hex/internal/mooc"
)

type CourseRepository struct {
	db      *sql.DB
	timeout time.Duration
}

func NewCourseRepository(db *sql.DB, timeout time.Duration) *CourseRepository {
	return &CourseRepository{
		db:      db,
		timeout: timeout,
	}
}

var courseStruct = sqlbuilder.NewStruct(new(sqlCourse))

func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	query, args := courseStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().Value(),
		Name:     course.Name().Value(),
		Duration: course.Duration().Value(),
	}).Build()

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

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

var NotFoundError = errors.New("could not the a course")

func (r *CourseRepository) Find(ctx context.Context, id mooc.CourseID) (mooc.Course, error) {
	qb := sqlbuilder.Select("*").From(sqlCourseTable)
	qb.Where(qb.Equal("id", id.Value()))
	query, args := qb.Build()

	rows, err := r.db.QueryContext(ctx, query, args...)

	if rows.Next() == false {
		return mooc.Course{}, fmt.Errorf("%w with id: %s", NotFoundError, id.Value())
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
