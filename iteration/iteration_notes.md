# Iteration
To do stuff repeatedly in go we only need ```for```. In go there is no
```while, do, until``` keywords.
```bash
func Repeat(character string) string {
	var repeatedString string
	for i := 0; i < repeatCount; i++ {
		repeatedString += character
	}
	return repeatedString
}
```

* by default benchmarks are run sequentially.