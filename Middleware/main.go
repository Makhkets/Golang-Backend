package main

import (
	"fmt"
	"log"
	"net/http"
)

func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("~~ ADMIN MIDDLEWARE DETECTED ~~")
		next.ServeHTTP(w, r)
	})
}

func userMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("~~ USER MIDDLEWARE DETECTED ~~")
		next.ServeHTTP(w, r)
	})
}

func allMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("~~ ALL MIDDLEWARE DETECTED ~~")
		next.ServeHTTP(w, r)
	})
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "TEST HANDLER")
}

func main() {
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/test/", testHandler)
	adminMux.HandleFunc("/admin/test2/", testHandler)

	// set middleware (adminMux)
	adminMW := adminMiddleware(adminMux)

	userMux := http.NewServeMux()
	userMux.HandleFunc("/user/test/", testHandler)
	userMux.HandleFunc("/user/test2/", testHandler)

	// set middleware (userMux)
	userMW := userMiddleware(userMux)

	allMux := http.NewServeMux()

	allMux.Handle("/admin/", adminMW)
	allMux.Handle("/user/", userMW)

	allMux.HandleFunc("/all/test/", testHandler)
	allMux.HandleFunc("/all/test2/", testHandler)

	log.Println("Started on address: http://localhost:8000")
	http.ListenAndServe(":8000", allMux)
}
