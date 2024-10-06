package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nabin3/LearnGoWithTests/dependency_injection"
	"github.com/nabin3/LearnGoWithTests/helloWorld"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	dependency_injection.Greet(w, "world")
}

func main() {
	sayHello, _ := helloWorld.Hello("Spanish", "Nolan")
	fmt.Println(sayHello)

	dependency_injection.Greet(os.Stdout, "Chris")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
