* map[comparable_type]any_type.
* An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap). This may make them feel like a "reference type", but they are not: ```A map value is a pointer to a runtime.hmap structure.```
* So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.
* Map can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.So, you should never initiallize a nil map variable:```var m map[string]string```. Instaed, you can initialize an empty map or use the make keyword to create amap for you
```bash
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```

