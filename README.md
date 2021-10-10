go-unidecode
==============

ASCII transliterations of Unicode text.


Installation
------------

```
go get -u github.com/lemontreeran/go-unidecode
```

install CLI tool:

```
go get -u github.com/lemontreeran/go-unidecode/unidecode
$ unidecode 北京
Bei Jing 
```


Usage
------

```go
package main
import (
	"fmt"
	"github.com/lemontreeran/go-unidecode"
)
func main() {
	s = "abc"
	fmt.Println(unidecode.Unidecode(s))
	// Output: abc
	s = "北京"
	fmt.Println(unidecode.Unidecode(s))
	// Output: Bei Jing
	s = "30 𝗄𝗆/𝗁"
	fmt.Println(unidecode.Unidecode(s))
	// Output: 30 km/h
	s = "kožušček"
	fmt.Println(unidecode.Unidecode(s))
	// Output: kozuscek
}
```
