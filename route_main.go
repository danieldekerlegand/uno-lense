package main

import (
	"uno-lense/data"
	"net/http"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	rcs, err := data.RemoteConnections()
	if err != nil {
		error_message(writer, request, "Cannot get remote connections")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, nil, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, rcs, "layout", "private.navbar", "index")
		}
	}
	// lessons, err := data.Lessons()
	// if err != nil {
	// 	error_message(writer, request, "Cannot get lessons")
	// } else {
	// 	_, err := session(writer, request)
	// 	if err != nil {
	// 		generateHTML(writer, lessons, "layout", "public.navbar", "index")
	// 	} else {
	// 		generateHTML(writer, lessons, "layout", "private.navbar", "index")
	// 	}
	// }
}
