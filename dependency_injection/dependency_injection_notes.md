```bash
func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}
```
Calling fmt.Printf prints to stdout, which is pretty hard for us to capture using the testing framework. What we need to do is to be able to inject (which is just a fancy word for pass in) the dependency of printing. Our function doesn't need to care how the printing happens, so we should accept an interface rather than a concrete type.
If we look at the source code of ```fmt.Printf``` we can see a way for us to hook in
```bash
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}
```
Interesting! Under the hood Printf just calls Fprintf passing in os.Stdout.
```bash
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```
An ```io.Writer```
```bash
type Writer interface {
	Write(p []byte) (n int, err error)
}
```
From this we can infer that ```os.Stdout``` implements ```io.Writer```; ```Printf``` passes ```os.Stdout``` to ```Fprintf``` which expects an ```io.Writer```.

```bash
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```
The Buffer type from the bytes package implements the Writer interface, because it has the method Write(p []byte) (n int, err error).

So we'll use it in our test to send in as our Writer and then we can check what was written to it after we invoke Greet