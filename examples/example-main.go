package main

import (
	"fmt"

	J "github.com/lacikawiz/go-json-util"
)

func main() {
	t1 := J.FromJSON(`[{"x":1},{"z":"str"},{"y":true}]`)
	//converting an object or Map to String produces a concise printout of the contents
	fmt.Println("Content of t1:", t1.String(nil)) //nil means we don't need error handling

	//accessing the value of t1[0]["x"] converted to Int
	{
		t2 := t1.I(0).K("x").Int(nil) //nil means we don't care about error handling
		fmt.Println("t2 = t1[0]['x'] : ", t2)
	}

	//accessing a non-existent key
	{
		var err error

		t3 := t1.I(0).K("a").Int(&err)

		fmt.Println("t3 = t1[0]['a'] : ", t3, "error=", err)
	}
}
