# Convert Utils

## Cast to primitive

support

int, int8, int16, int32, int64, float32, float64, bool, string
```go 
import (
    "github.com/optim-kazuhiro-seida/go-advance-type/convert"
)

func main() {
	a := convert.MustStr(120)             // var a string = "120" : Same SafeStr(120, "")
	b := convert.MustInt("120")           // var b int = 120      : Same SafeInt("120", 0)
	c := convert.MustInt64([]byte("120")) // var c int64 = 120    : Same SafeInt64([]byte("120"), 0)

	// With default value
	d := convert.SafeInt32(struct{}{}, 21)    // var d int32  = 21
	e := convert.SafeFloat32([]byte(""), 1.2) // var e float32 = 1.2

	// If you wants to handle err
	f, err := convert.Float64("hogehoge") // error
	g, err := convert.Bool("gegege")      // error

}
```

## Convert struct and map

```go
import (
    "github.com/optim-kazuhiro-seida/go-advance-type/convert"
)

func main() {
    a := Hoge {
        a: "a",
    }
    convert.Struct2Map(Hoge) // map[string]interface{}{"a": "a"}
}
```

## Exchange to ptr variable

```go
import (
	"github.com/optim-kazuhiro-seida/go-advance-type/convert"
)

func main() {
	ptrString := convert.StringPtr("test") // &"test"
	ptrInt := convert.IntPtr(2)            // &2
}
```

