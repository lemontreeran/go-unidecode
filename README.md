go-unidecode
==============
[![Build Status](https://travis-ci.org/github/lemontreeran/go-unidecode.svg?branch=master)](https://travis-ci.com/github/lemontreeran/go-unidecode)
[![Coverage Status](https://coveralls.io/repos/lemontreeran/go-unidecode/badge.svg?branch=master)](https://coveralls.io/r/lemontreeran/go-unidecode?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/lemontreeran/go-unidecode)](https://goreportcard.com/report/github.com/lemontreeran/go-unidecode)
[![GoDoc](https://godoc.org/github.com/lemontreeran/go-unidecode?status.svg)](https://godoc.org/github.com/lemontreeran/go-unidecode)

ASCII transliterations of Unicode text.


Installation
------------

```
go get -u github.com/lemontreeran/go-unidecode
```

install CLI tool:

```
go get -u github.com/lemontreeran/go-unidecode/unidecode

$ unidecode 北京kožušček
Bei Jing kozuscek
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
