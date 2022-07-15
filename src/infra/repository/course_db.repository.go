package repository

import (
	"database/sql"
	"project1/src/entity"
)

type CourseMySQLRepository struct {
	Db *sql.DB
}

func (c *CourseMySQLRepository) Insert(course entity.Course) error {
	stmp, err := c.Db.Prepare(`Insert into courses(id, name, description, status) values(?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmp.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status,
	)

	if err != nil {
		return err
	}

	return nil
}
