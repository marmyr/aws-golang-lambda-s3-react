package main // Important: Packages with endpoints must be named 'main'
import (
	"example.com/someservice555/endpoints"
	"example.com/someservice555/internal/routing"
	"log"
)

func main() {
	engine := routing.Build()

	for _, e := range endpoints.Endpoints {
		routing.AddRoute(engine, e.Path, e.Method, e.Function)
	}

	if err := engine.Run(); err != nil {
		log.Printf("Error starting gin %v", err)
	}

	log.Printf("Application exiting.")
}