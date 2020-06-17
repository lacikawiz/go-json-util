package main

import (
	"fmt"

	J "github.com/lacikawiz/go-json-util"
)

func main() {
	t1 := J.FromJSON([]byte(`[{"x":1},{"z":"str"},{"y":true}]`))
	fmt.Printf("%+v", t1)

	//accessing the value of t1[0]["x"] converted to Int
	t2 := t1.I(0).K("x").Int()
	fmt.Println("t2=", t2)

	//accessing a non-existent key
	{
		x := t1.I(0).K("a")
		t3 := x.Int()
		fmt.Println("t3=", t3, "error=", x.Err)
	}
}
