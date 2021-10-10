package unidecode_test

import (
	"fmt"

	"github.com/lemontreeran/go-unidecode"
)

func ExampleUnidecode() {
	s := "北京kožušček"
	fmt.Println(unidecode.Unidecode(s))
	// Output: Bei Jing kozuscek
}

func ExampleUnidecode_ascii() {
	s := "abc"
	fmt.Println(unidecode.Unidecode(s))
	// Output: abc
}
