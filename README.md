# go-sqlite4-varuint

Go package implements SQLite4-like variable unsigned integer encoding
as described in http://www.sqlite.org/src4/doc/trunk/www/varint.wiki.

Unlike varint from encoding/binary package, this encoding uses fewer bytes
for smaller values, and the number of encoded bytes can be determined by
looking at the first byte.

### Install
   
`go get github.com/encobrain/go-sqlite4-varuint`

### Usage

```go
package main

import (
	"fmt"
	"github.com/encobrain/go-sqlite4-varuint"
)

func main() {
	buf := make([]byte, varuint.MaxBufSize)
	
	val := uint64(12345)
	
	s := varuint.EncodeSize(val)
	fmt.Println(s)
	// Output: 3

	n := varuint.Encode(buf, 12345)

	if n != s {
		panic(fmt.Errorf("Written size  != Encode size"))
    }
	
	buf = buf[:n]

	fmt.Printf("% x\n", buf)
	// Output: f9 27 49
	
	isDecodable := varuint.IsDecodable(buf[:2])
	fmt.Printf("%v\n", isDecodable)
	// Output: false

	v, n := varuint.Decode(buf)
	fmt.Println(v, n)
	// Output: 12345, 3
}

```