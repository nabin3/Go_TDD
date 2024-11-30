package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nabin3/LearnGoWithTests/dependency_injection"
	"github.com/nabin3/LearnGoWithTests/helloWorld"
	"github.com/nabin3/LearnGoWithTests/mocking"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	dependency_injection.Greet(w, "world")
}

func main() {
	sayHello, _ := helloWorld.Hello("Spanish", "Nolan")
	fmt.Println(sayHello)

	// To see mocking code in action
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, Csleep: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
	fmt.Println()

	// To see dependency injection code in action
	dependency_injection.Greet(os.Stdout, "Chris")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
