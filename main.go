package main

import (
	"net/http"
	"time"
)

func main() {
	p("Lense", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_lesson.go
	mux.HandleFunc("/lesson/new", newLesson)
	mux.HandleFunc("/lesson/create", createLesson)
	mux.HandleFunc("/lesson/component", componentLesson)
	mux.HandleFunc("/lesson/read", readLesson)
	mux.HandleFunc("/lesson/publish", publishLesson)
	mux.HandleFunc("/lesson/unpublish", unpublishLesson)

	mux.HandleFunc("/local/images", localImages)

	mux.HandleFunc("/repository", repository)
	mux.HandleFunc("/repository/images", repositoryImages)
	mux.HandleFunc("/repository/push", repositoryPush)
	mux.HandleFunc("/repository/pull", repositoryPull)

	mux.HandleFunc("/settings", settings)
	mux.HandleFunc("/connect", connect)
	mux.HandleFunc("/disconnect", disconnect)

	mux.HandleFunc("/container/stop", containerStop);
	mux.HandleFunc("/container/pause", containerPause);
	mux.HandleFunc("/container/start", containerStart);
	mux.HandleFunc("/container/restart", containerRestart);
	mux.HandleFunc("/container/download", containerDownload);

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
