* We can create a simple type using a struct. A struct is just a named collection of fields where you can store data.

* When we need to calculate area of rectangle and circle both in other language we can do 
```bash
func Area(circle Circle) float64       {}
func Area(rectangle Rectangle) float64 {}
```
but in go it's not allowed. So we can solve the problem by writting area method for circle in different package but that will be unefficient. The efficient way is to use method.

* A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

* Methods are very similar to functions but they are called by invoking them on an instance of a particular type. 

* It's convention to use first letter of type as reciever name in method, here the reciever will always be in lowercase letter.

* Interfaces are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety.