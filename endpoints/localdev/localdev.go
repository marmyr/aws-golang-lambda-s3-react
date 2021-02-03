package main // Important: Packages with endpoints must be named 'main'
import (
	"example.com/someservice555/api/hello"
	"example.com/someservice555/api/welcome"
	"example.com/someservice555/internal/routing"
	"log"
)

func main() {
	engine := routing.Build()
	routing.AddRoute(engine, hello.Path, hello.Method, hello.ProcessRequest)
	routing.AddRoute(engine, welcome.Path, welcome.Method, welcome.ProcessRequest)

	if err := engine.Run(); err != nil {
		log.Printf("Error starting gin %v", err)
	}

	log.Printf("Application exiting.")
}