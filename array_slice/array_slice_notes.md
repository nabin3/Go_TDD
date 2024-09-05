# Arrays and Slices
array allows us to store multiple elements of the same type in a variable in a particular order. Arrays have a fixed capacity which we
define when we declare the variable. We can initialize an array in two ways:
[N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
[...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5} 

* An interesting property of arrays is that the size is encoded in its type. If you try to pass an [4]int into a function that expects [5]int, it won't compile. They are different types so it's just the same as trying to pass a string into a function that wants an int.

* Go has **slice** which doesn't encode size of the collection and it can have any size. Example: ```mySlice := []int{1,2,3}```

* Whilst 100% coverage should not be our end goal, still it is useful to use ```go test -cover``` to identify areas of our code not covered by tests.