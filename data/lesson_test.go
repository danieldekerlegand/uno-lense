package data


import (
	"testing"
)

// Delete all lessons from database
// func LessonDeleteAll() (err error) {
// 	db := db()
// 	defer db.Close()
// 	statement := "delete from lessons"
// 	_, err = db.Exec(statement)
// 	if err != nil {
// 		return
// 	}
// 	return
// }
//
// func Test_CreateLesson(t *testing.T) {
// 	setup()
// 	if err := users[0].Create(); err != nil {
// 		t.Error(err, "Cannot create user.")
// 	}
// 	conv, err := users[0].CreateLesson("My first lesson")
// 	if err != nil {
// 		t.Error(err, "Cannot create lesson")
// 	}
// 	if conv.UserId != users[0].Id {
// 		t.Error("User not linked with lesson")
// 	}
// }
