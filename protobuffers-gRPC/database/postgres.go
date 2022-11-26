package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/jamsmendez/protobuffers-gRPC/models"
	"golang.org/x/net/context"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, err
}

func (p *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := p.db.ExecContext(
		ctx,
		"INSERT INTO students (id, name, age) VALUES ($1, $2, $3)",
		student.ID,
		student.Name,
		student.Age,
	)

	return err
}

func (p *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("Postgrest.GetStudent.Rows.Close: ", err)
		}
	}(rows)

	var students = []models.Student{}

	for rows.Next() {
		student := models.Student{}

		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
		)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	if len(students) == 0 {
		return nil, nil
	}

	student := students[0]

	return &student, nil
}

func (p *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := p.db.ExecContext(
		ctx,
		"INSERT INTO tests (id, name) VALUES ($1, $2)",
		test.ID,
		test.Name,
	)

	return err
}

func (p *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("Postgrest.GetTest.Rows.Close: ", err)
		}
	}(rows)

	var tests = []models.Test{}

	for rows.Next() {
		test := models.Test{}

		err := rows.Scan(
			&test.ID,
			&test.Name,
		)

		if err != nil {
			return nil, err
		}

		tests = append(tests, test)
	}

	if len(tests) == 0 {
		return nil, nil
	}

	test := tests[0]

	return &test, nil
}

func (p *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := p.db.ExecContext(
		ctx,
		"INSERT INTO questions (id, question, answer, test_id) VALUES ($1, $2, $3, $4)",
		question.ID,
		question.Question,
		question.Answer,
		question.TestID,
	)

	return err
}

func (repo *PostgresRepository) SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO enrollments(student_id, test_id) VALUES($1, $2)",
		enrollment.StudentID, enrollment.TestID,
	)

	return err
}

func (repo *PostgresRepository) GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT id, name, age FROM students WHERE id IN (SELECT student_id FROM enrollments WHERE test_id = $1)",
		testId,
	)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var students []*models.Student

	for rows.Next() {
		student := models.Student{}
		if err = rows.Scan(&student.ID, &student.Name, &student.Age); err == nil {
			students = append(students, &student)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (repo *PostgresRepository) GetQuestionsPerTest(ctx context.Context, testID string) ([]*models.Question, error) {
	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT id, question FROM questions WHERE test_id = $1",
		testID,
	)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var questions []*models.Question

	for rows.Next() {
		question := models.Question{}
		if err = rows.Scan(&question.ID, &question.Question); err == nil {
			questions = append(questions, &question)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}
