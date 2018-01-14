package main

import (
	"fmt"
	"uno-lense/data"
	"net/http"
)

// GET /lesson/new
// Show the new lesson form page
func newLesson(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.lesson")
	}
}

// POST /lesson/new
// Create a new lesson
func createLesson(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		// user, err := sess.User()
		// if err != nil {
		// 	danger(err, "Cannot get user from session")
		// }
		base_image := request.PostFormValue("image")
		data.PushImage(base_image)
		// topic := request.PostFormValue("topic")
		// base_image := request.PostFormValue("image")
		// details := request.PostFormValue("details")
		// fmt.Println("base_image " + base_image)
		// if _, err := user.CreateLesson(topic, base_image, details); err != nil {
		// 	danger(err, "Cannot create lesson")
		// }
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

// POST /local/images
func localImages(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/", 302)
	} else {
		var images []byte
		images, err = data.ListImages()

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(images)
	}
}

// GET /repository
func repository(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "repository")
	}
}

// POST /repository/images
func repositoryImages(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		repo := request.PostFormValue("repo")
		username := request.PostFormValue("username")
		password := request.PostFormValue("password")

		var images []byte
		images, err = data.ListRemoteImages(repo, username, password)

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(images)
	}
}

// POST
func repositoryPush(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		repo := request.PostFormValue("image")
		data.PushImage(repo)

		http.Redirect(writer, request, "repository", 302)
	}
}

// POST
func repositoryPull(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		repo := request.PostFormValue("image")
		data.PullImage(repo)

		http.Redirect(writer, request, "repository", 302)
	}
}

// POST
// /connect
func connect(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	ip := request.RemoteAddr
	name := request.PostFormValue("name")
	s_id := request.PostFormValue("s_id")

	fmt.Println("ip", ip)
	fmt.Println("name", name)
	fmt.Println("s_id", s_id)
}

// POST
// /disconnect
func disconnect(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	ip := request.RemoteAddr
	name := request.PostFormValue("name")
	s_id := request.PostFormValue("s_id")

	fmt.Println("ip", ip)
	fmt.Println("name", name)
	fmt.Println("s_id", s_id)
}
