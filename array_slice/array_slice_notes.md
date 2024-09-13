# Arrays and Slices
array allows us to store multiple elements of the same type in a variable in a particular order. Arrays have a fixed capacity which we
define when we declare the variable. We can initialize an array in two ways:
[N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
[...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5} 

* An interesting property of arrays is that the size is encoded in its type. If you try to pass an [4]int into a function that expects [5]
int, it won't compile. They are different types so it's just the same as trying to pass a string into a function that wants an int.

* Go has **slice** which doesn't encode size of the collection and it can have any size. Example: ```mySlice := []int{1,2,3}```

* Whilst 100% coverage should not be our end goal, still it is useful to use ```go test -cover``` to identify areas of our code not covered by tests.

* We can't simply compare two slices with ```== or !=```, we can do those to check if a slice is nil or not. We cam use ```DeepEqual```
from reflect package two see if two slices are equal or not(It is also useful for seeing if any two variables are the same.) But this func is
not typesafe i.e we can compare different types and the program still compiles. One thing we can use is ```Equal``` func from
```slices``` package. Keep in mind that this ```slices.Equal``` only work with comparable typed slice, so we **can not compare two 2D slices** with that. 

* We can create a slice with ```make``` function like this:
```bash
mySlice := make([]type, length, capacity)       
# length = number of items in the slice and
# capacity = number of items slice's underlying array can contain
```

* append function add a value in a slice, but when a slice reaches it's capacity and still we use append() to add a value to it then a new slice with larger capacity is created and all value of old slice will be copied to it and then the new value will be added, so to capture that new slice we can reassign the return value of append to the old slice variable.
```bash
old_slice = append(old_slice, value)
```

* We could've created a new function checkSums like we normally do, but in this case, we're showing a new technique, assigning a function to a variable. It might look strange but, it's no different to assigning a variable to a string, or an int, functions in effect are values too.

* It's not shown here, but this technique can be useful when you want to bind a function to other local variables in "scope" (e.g between some {}). It also allows you to reduce the surface area of your API.

* By defining this function inside the test, it cannot be used by other functions in this package. Hiding variables and functions that don't need to be exported is an important design consideration.

* A handy side-effect of this is this adds a little type-safety to our code. If a developer mistakenly adds a new test with checkSums(t, got, "dave") the compiler will stop them in their tracks.
```bash
$ go test
./sum_test.go:52:21: cannot use "dave" (type string) as type []int in argument to checkSums
```

* Create a slice from an array and changing that slice will mutate the original array.

* It's a goo idea to make copy of a slice, created from an large slice.
```bash
func main() {
	a := make([]int, 1e6) // slice "a" with len = 1 million
	b := a[:2] // even though "b" len = 2, it points to the same the underlying array "a" points to
	
	c := make([]int, len(b)) // create a copy of the slice so "a" can be garbage collected
	copy(c, b)
	fmt.Println(c)
}
```