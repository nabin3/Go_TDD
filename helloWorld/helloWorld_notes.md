* When we write a program in golang, we will have a main package defined with a ```main``` func inside it. Packages are ways of grouping up related go code together. There can only one main function.
* It is good idea to separate **domain** code from side effect. The ```fmt.Println``` is also side effect(printing to stdout).
```bash
package main

import "fmt"

// "Hello, World" is our domain code, which has kept separate from fmt.Println() side effect
func Hello() string {
	return "Hello, World"
}

func main() {
	fmt.Println(Hello())
}
```
* A test filename in go should be end wiyh **_test** and the test func name must be begin with **Test**.
* From **go 1.16** or later go test give error with out existence of an **go.mod** file
* ```go mod init module_name or go mod init github.com/github_username/module_name``` to create a go.mod file. **go.mod** file gives the location where the module can be obtained and it list dependencies of the module.
* With ```Run``` function we can group related tests in a single **Test** function.
* A ```Helper``` function is used for checking for the output of ```Hello``` func here. Helper func's speciality is if any thing wrong in our code it will report the error in calling place of it not inside of it.
* can install godoc with this ```go install golang.org/x/tools/cmd/godoc@latest``` for locally view godoc and code example mentioned in ```Example func like here ExampleHello func``` this Example func should be in ```_test.go``` file. See I have include ```//Output: desired output message``` in Example func. If don't included ```// Output: ``` then this Example function will be compiled but wiull not execute when we run test. run ``` go test -v``` to see test result in detail.
```bash
func ExampleHello() {
	fmt.Println(Hello("Spanish", "Nabin"))

	// Output: Hola, Nabin
}
```

* ```Benchmark``` func tion will check efficiency of a function. This func should be kept inside ```_test.go``` file and to run benchmark test run this ``` go test -bench=. ```
```bash
func BenchmarkHello(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Hello("Spanish", "Nabin")
	}
}
```