# genex [![GoDoc](https://godoc.org/github.com/alixaxel/genex?status.svg)](https://godoc.org/github.com/alixaxel/genex) [![GoCover](http://gocover.io/_badge/github.com/alixaxel/genex)](http://gocover.io/github.com/alixaxel/genex) [![Go Report Card](https://goreportcard.com/badge/github.com/alixaxel/genex)](https://goreportcard.com/report/github.com/alixaxel/genex)

Genex package for Go

Easy and efficient package to expand any given regex into all the possible strings that it can match.

This is the code that powers [namegrep](https://namegrep.com/).

## Usage

### With options struct

```go
package main

import (
    "fmt"
    "regexp/syntax"

    //"github.com/alixaxel/genex"
    "github.com/fyxme/genex"
)

/* Default values
	DEFAULT_CHARSET = "[a-zA-Z0-9]" // charset that will be used when arbytrary values are specified (ie. a dot (.))
	DEFAULT_INFINITE = 3 // number of interations for a multiplier (ie. a + or * multiplier)
*/

func main() {
    charset, _ := syntax.Parse(`[0-9a-z]`, syntax.Perl)

    // options can also be nil when passed to array if you are happy with the default options
    options := genex.GenOpts {
        charset: charset, // set to nil if you are ok with the default charset
        infinite: 5, 
    }   

    if input, err := syntax.Parse(`(foo|bar|baz){1,2}\d`, syntax.Perl); err == nil {
    	fmt.Println("Count:", genex.Count(input, options))

    	genex.Generate(input, options, func(output string) {
    		fmt.Println("[*]", output)
    	})
    }
}

```

### Without options struct 

```go
package main

import (
    "fmt"
    "regexp/syntax"

    //"github.com/alixaxel/genex"
    "github.com/fyxme/genex"
)

func main() {
    charset, _ := syntax.Parse(`[0-9a-z]`, syntax.Perl)

    if input, err := syntax.Parse(`(foo|bar|baz){1,2}\d`, syntax.Perl); err == nil {
    	fmt.Println("Count:", genex.CountWithoutOptions(input, charset, 3))

    	genex.GenerateWithoutOptions(input, charset, 3, func(output string) {
    		fmt.Println("[*]", output)
    	})
    }
}
```

## Output

```
Count: 120

[*] foo0
[*] ...
[*] foo9
[*] foofoo0
[*] ...
[*] foofoo9
[*] foobar0
[*] ...
[*] foobar9
[*] foobaz0
[*] ...
[*] foobaz9
[*] bar0
[*] ...
[*] bar9
[*] barfoo0
[*] ...
[*] barfoo9
[*] barbar0
[*] ...
[*] barbar9
[*] barbaz0
[*] ...
[*] barbaz9
[*] baz0
[*] ...
[*] baz9
[*] bazfoo0
[*] ...
[*] bazfoo9
[*] bazbar0
[*] ...
[*] bazbar9
[*] bazbaz0
[*] ...
[*] bazbaz9
```

## Install

	go get github.com/alixaxel/genex

## License

MIT
