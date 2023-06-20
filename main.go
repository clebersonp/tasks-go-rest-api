package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/clebersonp/tasks-go-rest-api/configs"
	"github.com/clebersonp/tasks-go-rest-api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	// config the log formater
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	// create a router. https://github.com/go-chi/chi
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rest routes for "tasks" resource
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", handlers.List)    // GET /tasks
		r.Post("/", handlers.Create) // POST /tasks

		// subrouters:
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.Get)       // GET /tasks/123
			r.Put("/", handlers.Update)    // PUT /tasks/123
			r.Delete("/", handlers.Delete) // DELETE /tasks/123
		})
	})

	// starts the rest api server
	serverPort := configs.ServerPort()
	log.Println("Started the rest api server at port:", serverPort)
	http.ListenAndServe(fmt.Sprintf("localhost:%s", serverPort), r)
}
