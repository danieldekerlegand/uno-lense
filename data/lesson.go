package data

import (
	"time"
	"fmt"
)

type Lesson struct {
	Id        int
	Uuid      string
	Topic     string
	Details		string
	Published int
	UserId    int
	CreatedAt time.Time
	BaseImage string
}

type Component struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	LessonId  int
	CreatedAt time.Time
}

// format the CreatedAt date to display nicely on the screen
func (lesson *Lesson) CreatedAtDate() string {
	return lesson.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (component *Component) CreatedAtDate() string {
	return component.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of components in a lesson
func (lesson *Lesson) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM components where lesson_id = $1", lesson.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// get components to a lesson
func (lesson *Lesson) Components() (components []Component, err error) {
	rows, err := Db.Query("SELECT id, uuid, body, user_id, lesson_id, created_at FROM components where lesson_id = $1", lesson.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		component := Component{}
		if err = rows.Scan(&component.Id, &component.Uuid, &component.Body, &component.UserId, &component.LessonId, &component.CreatedAt); err != nil {
			return
		}
		components = append(components, component)
	}
	rows.Close()
	return
}

// Create a new lesson
func (user *User) CreateLesson(topic string, base_image string, details string) (conv Lesson, err error) {
	statement := "insert into lessons (uuid, topic, user_id, created_at, base_image, details, published) values ($1, $2, $3, $4, $5, $6, 0)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// // use QueryRow to return a row and scan the returned id into the Session struct
	// err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	// return

	uuid := createUUID()
	// execute the insert
	_, err = stmt.Exec(uuid, topic, user.Id, time.Now(), base_image, details)

	// scan the new into the Session struct
	statement = "SELECT id, uuid, topic, user_id, created_at, base_image, details, published FROM lessons WHERE uuid = $1"
	stmt, err = Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(uuid).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt, &conv.BaseImage, &conv.Published)
	return
}

// Create a new component to a lesson
func (user *User) CreateComponent(conv Lesson, body string) (component Component, err error) {
	statement := "insert into components (uuid, body, user_id, lesson_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, lesson_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&component.Id, &component.Uuid, &component.Body, &component.UserId, &component.LessonId, &component.CreatedAt)
	return
}

// Get all lessons in the database and returns it
func Lessons() (lessons []Lesson, err error) {
	fmt.Println("in Lessons()")
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at, base_image, details, published FROM lessons ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Lesson{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt, &conv.BaseImage, &conv.Details, &conv.Published); err != nil {
			return
		}
		lessons = append(lessons, conv)
	}
	rows.Close()
	return
}

// Get a lesson by the UUID
func LessonByUUID(uuid string) (conv Lesson, err error) {
	conv = Lesson{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at, base_image, details, published FROM lessons WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt, &conv.BaseImage, &conv.Details, &conv.Published)
	return
}

// Get the user who started this lesson
func (lesson *Lesson) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", lesson.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// Get the user who wrote the component
func (component *Component) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", component.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func PublishLesson(uuid string) {
	statement := "UPDATE lessons SET published = 1 WHERE uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
}

func UnpublishLesson(uuid string) {
	statement := "UPDATE lessons SET published = 0 WHERE uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
}
