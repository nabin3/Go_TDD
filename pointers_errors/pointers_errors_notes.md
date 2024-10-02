* In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.

* Pointers to struct also known as struct pointers and they are automatically dereferenced.

* The syntax for creating a method on a type declaration is the same as it is on a struct.

* the **fmt** pkg has an interface name ```stringer``` that lets us define how our type is ptinted when used with the ```%s``` format string in prints.
```bash
type Stringer interface {
	String() string
}
```

* nil is synonymous with null from other programming languages. Errors can be nil because the return type of Withdraw will be error, which is an interface. If you see a function that takes arguments or returns values that are interfaces, they can be nillable. Like null if you try to access a value that is nil it will throw a runtime panic. This is bad! You should make sure that you check for nils.

* Errors can be converted to a string with the .Error() method

* ```go install github.com/kisielk/errcheck@latest``` will install ```errcheck```. Run ```errcheck``` inside the directory with your code 
to check for potential error which didn't covered with your test.

# Wrapping Up:

## Pointers:
* Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
* The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

## nil:
* Pointers can be nil
* When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.
* Useful for when you want to describe a value that could be missing

## Errors:
* Errors are the way to signify failure when calling a function/method.
* By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.
* Don’t just check errors, handle them gracefully. See this study: ```https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully```

## Create new types from existing ones
*Useful for adding more domain specific meaning to values
* Can let you implement interfaces