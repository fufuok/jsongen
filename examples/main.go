package main

import (
	"fmt"

	"github.com/fufuok/bytespool"
	"github.com/fufuok/jsongen"
)

func main() {
	js := jsongen.NewMap()
	js.PutString("s", `a"b"\c`)
	js.PutFloat("f", 3.14)
	js.PutBool("b", false)
	jsArr := jsongen.NewArray()
	jsArr.AppendInt(7)
	jsArr.AppendStringArray([]string{"A", "B"})
	js.PutArray("sub", jsArr)
	js.PutRawString("raw", `{"n":null,"m":[1,"ff"]}`)

	size := js.Size()
	bs := bytespool.Get(size)
	defer bytespool.Put(bs)
	data := js.Serialize(bs[:0])

	// 也可以直接使用 nil
	// data := js.Serialize(nil)

	fmt.Printf("%s\n", data)

	// Output:
	// {"s":"a\"b\"\\c","f":3.14,"b":false,"sub":[7,["A","B"]],"raw":{"n":null,"m":[1,"ff"]}}
}
