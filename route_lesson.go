package main

import (
	"fmt"
	"uno-lense/data"
	"net/http"
)

// GET /lessons/new
// Show the new lesson form page
func newLesson(writer http.ResponseWriter, request *http.Request) {
	// listImages()
	// listRunningContainers()
	images := data.Images()
	fmt.Printf("images: %+v \n", images)
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, images, "layout", "private.navbar", "new.lesson")
	}
}

// POST /signup
// Create the user account
func createLesson(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		base_image := request.PostFormValue("image")
		details := request.PostFormValue("details")
		fmt.Println("base_image " + base_image)
		if _, err := user.CreateLesson(topic, base_image, details); err != nil {
			danger(err, "Cannot create lesson")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /lesson/read
// Show the details of the lesson, including the components and the form to write a component
func readLesson(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	lesson, err := data.LessonByUUID(uuid)
	if err != nil {
		error_message(writer, request, "Cannot read lesson")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &lesson, "layout", "public.navbar", "public.lesson")
		} else {
			generateHTML(writer, &lesson, "layout", "private.navbar", "private.lesson")
		}
	}
}

func publishLesson(writer http.ResponseWriter, request *http.Request) {
	uuid := request.PostFormValue("uuid")

	data.PublishLesson(uuid)
	url := fmt.Sprint("/lesson/read?id=", uuid)
	http.Redirect(writer, request, url, 302)
}

func unpublishLesson(writer http.ResponseWriter, request *http.Request) {
	uuid := request.PostFormValue("uuid")

	data.UnpublishLesson(uuid)
	url := fmt.Sprint("/lesson/read?id=", uuid)
	http.Redirect(writer, request, url, 302)
}

// POST /lesson/component
// Create the component
func componentLesson(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		lesson, err := data.LessonByUUID(uuid)
		if err != nil {
			error_message(writer, request, "Cannot read lesson")
		}
		if _, err := user.CreateComponent(lesson, body); err != nil {
			danger(err, "Cannot create component")
		}
		url := fmt.Sprint("/lesson/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
