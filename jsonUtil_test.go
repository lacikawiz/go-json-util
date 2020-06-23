package jsonutil

import (
	"errors"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Run("Array with objects", func(t *testing.T) {
		t1 := FromJSON(`[{"x":1},{"z":"str"},{"y":true},null]`)
		if t1.I(0).K("x").Int(nil) != 1 {
			t.Error("int test failed")
		}
		if t1.I(1).K("z").String(nil) != "str" {
			t.Error("string test failed")
		}
		if !t1.I(2).K("y").Bool(nil) {
			t.Error("bool test failed")
		}
		if !t1.I(3).IsNil() {
			t.Error("null entry fail")
		}
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		t2 := FromJSON(`{x:1}`)
		if t2.Err == nil {
			t.Error("invalid JSON test failed")
		}
		t2 = FromJSON(`[w]`)
		if t2.Err == nil {
			t.Error("invalid JSON test failed")
		}
	})

	t3 := Obj{}
	t.Run("Nested Objects", func(t *testing.T) {
		t3 = FromJSON(`{"x":{
			"a":2,"b":"str","c":[1,2,"3"]
		}}`)
		if t3.K("x", "c").I(2).String(nil) != "3" {
			t.Error("nested JSON test failed")
		}
	})

	t.Run("Invalid Indexes", func(t *testing.T) {
		t4 := t3.K("x", "C")
		if !t4.IsNil() {
			t.Error("JSON walk test failed")
		}
		t4 = t3.K("x", "C").I(0)
		if t4.Err == nil {
			t.Error("JSON walk test failed")
		}
	})

	t.Run("wrong type of access", func(t *testing.T) {
		// t.Log("Testing wrong type of access(array when object is present)")
		t4 := t3.I(0)
		if t4.Err == nil {
			t.Error("JSON wrong type access test failed")
		}
	})
}

func TestEncode(t *testing.T) {
	t.Log("Testing encoding into JSON.")
	testStr := `[{"x":1},{"z":"str"},{"y":true}]`
	t1 := FromJSON(testStr)
	res := t1.ToJSON()
	if res != testStr {
		t.Error("JSON encoding test failed:")
		t.Log("Should have:", testStr)
		t.Log("Got", res)
	}
}

func TestAddKV(t *testing.T) {
	t1 := FromJSON(`[{"x":1},{"z":"str"},{"y":true}]`)
	t2 := t1.I(0)
	t2.AddKV("long", "wait")
	if t2.Err != nil {
		t.Error("Adding Key Value failed")
	}

	t2.AddKV("short", -10)
	if t2.Err != nil {
		t.Error("Adding Key Value failed")
	}

	if t1.I(0).K("long").String(nil) != "wait" ||
		t1.I(0).K("short").Int(nil) != -10 {
		t.Error("Adding Key Value failed")
	}
}

func TestSetError(t *testing.T) {
	var err error
	//should not crash
	setError(nil, errors.New("test"))
	setError(&err, errors.New("test"))
	if err.Error() != "test" {
		t.Error("setError() malfunction")
	}
	setError(&err, nil)
	if err != nil {
		t.Error("setError() malfunction")
	}
}

func TestConversions(t *testing.T) {
	t1 := FromJSON(`{"null":null,"num":-999999,"numStr":"1000","str":"hello","arr":[1,2],"obj":{"a":1},"bool":true,"boolStr1":"true","boolStr2":"yes"}`)
	var err error

	egNil := t1.K("null")
	egNum := t1.K("num")
	egNumStr := t1.K("numStr")
	egStr := t1.K("str")
	egArr := t1.K("arr")
	egObj := t1.K("obj")

	{ //String conversion
		egNil.String(nil) //should not crash
		{
			x := egNil.String(&err)
			if x != "" || err == nil {
				t.Error("`nil` to String conversion failed:")
			}
		}
		{
			x := egNum.String(&err)
			if x != "-999999" || err != nil {
				t.Error("num to String conversion failed:")
			}
		}
		{
			x := egStr.String(&err)
			if x != "hello" || err != nil {
				t.Error("string to String conversion failed:")
			}
		}
		{
			x := egArr.String(&err)
			if x != "[1 2]" || err != nil {
				t.Error("array to String conversion failed:")
			}
		}
		{
			x := egObj.String(&err)
			if x != "map[a:1]" || err != nil {
				t.Error("object to String conversion failed:")
			}
		}
	}
	{ //Int(64) conversion
		egNil.Int64(nil) //should not crash
		{
			x := egNil.Int64(&err)
			if x != 0 || err == nil {
				t.Error("`nil` to Int64 conversion failed:")
			}
		}
		{
			x := egNum.Int64(&err)
			if x != -999999 || err != nil {
				t.Error("num to Int64 conversion failed:")
			}
		}
		{
			x := egStr.Int64(&err)
			if x != 0 || err == nil {
				t.Error("string to Int64 conversion failed:")
			}
		}
		{
			x := egNumStr.Int64(&err)
			if x != 1000 || err != nil {
				t.Error("string to int64 conversion failed:")
			}
		}
		{
			x := egArr.Int64(&err)
			if x != 0 || err == nil {
				t.Error("array to int64 conversion failed:")
			}
		}
		{
			x := egObj.Int64(&err)
			if x != 0 || err == nil {
				t.Error("object to Int64 conversion failed:")
			}
		}
	}
	{ //Float64 conversion
		egNil.Float64(nil) //should not crash
		{
			x := egNil.Float64(&err)
			if x != 0 || err == nil {
				t.Error("`nil` to Float64 conversion failed:")
			}
		}
		{
			x := egNum.Float64(&err)
			if x != -999999 || err != nil {
				t.Error("num to Float64 conversion failed:")
			}
		}
		{
			x := egStr.Float64(&err)
			if x != 0 || err == nil {
				t.Error("string to Float64 conversion failed:")
			}
		}
		{
			x := egNumStr.Float64(&err)
			if x != 1000 || err != nil {
				t.Error("string to Float64 conversion failed:")
			}
		}
		{
			x := egArr.Float64(&err)
			if x != 0 || err == nil {
				t.Error("array to Float64 conversion failed:")
			}
		}
		{
			x := egObj.Float64(&err)
			if x != 0 || err == nil {
				t.Error("object to Float64 conversion failed:")
			}
		}
	}

	{ //Bool conversion
		egNil.Bool(nil) //should not crash
		{
			x := egNil.Bool(&err)
			if x != false || err != nil {
				t.Error("`nil` to Bool conversion failed:")
			}
		}
		{
			x := egNum.Bool(&err)
			if x != false || err == nil {
				t.Error("num to Bool conversion failed:")
			}
		}
		{
			x := egStr.Bool(&err)
			if x != false || err == nil {
				t.Error("string to Bool conversion failed:")
			}
		}

		{
			x := t1.K("boolStr1").Bool(&err)
			if x != true || err != nil {
				t.Error("string to Bool conversion failed:")
			}
		}
		{
			x := t1.K("boolStr2").Bool(&err)
			if x != true || err != nil {
				t.Error("string to Bool conversion failed:")
			}
		}

		{
			x := egArr.Bool(&err)
			if x != false || err == nil {
				t.Error("array to Bool conversion failed:")
			}
		}
		{
			x := egObj.Bool(&err)
			if x != false || err == nil {
				t.Error("object to Bool conversion failed:")
			}
		}
	}

	{ //ToStringArray
		t2 := FromJSON(`["a","b","c"]`)
		t3 := FromJSON(`[1,"a",true]`)
		{
			x := t2.ToStringArray(&err)
			if len(x) != 3 || x[0] != "a" || x[1] != "b" || x[2] != "c" || err != nil {
				t.Error("array to stringArray conversion failed")
			}
		}

		{
			x := t3.ToStringArray(&err)
			if len(x) != 3 || x[0] != "1" || x[1] != "a" || x[2] != "true" || err != nil {
				// fmt.Println("x=", x, "err=", err)
				t.Error("non-string array to stringArray conversion failed")
			}
		}
		{
			x := egNum.ToStringArray(&err)
			if len(x) != 0 || err == nil {
				t.Error("non-array to stringArray conversion failed")
			}
		}
	}
}

/*████████████████████████████████████████████
██████ ITERATORS
████████████████████████████████████████████*/

func TestIterators(t *testing.T) {
	t1 := FromJSON(`{"null":null,"num":-999999,"numStr":"1000","str":"hello","arr":[1,2],"obj":{"a":1,"b":2},"bool":true,"boolStr1":"true","boolStr2":"yes"}`)

	// var err error

	{ //ForEachArr
		if t1.K("null").ForEachArr(func(o Obj, idx int) {
			t.Error("iteration over null fail")
		}) != nil {
			t.Error("iteration over null fail")
		}

		count := 0
		if t1.K("arr").ForEachArr(func(o Obj, idx int) {
			count++
			if idx > 1 {
				t.Error("array iteration index fail")
			}
			if (idx == 0 && o.Int(nil) != 1) ||
				(idx == 1 && o.Int(nil) != 2) {
				t.Error("array iteration fail")
			}
		}) != nil {
			t.Error("iteration over array fail")
		}
		if count != 2 {
			t.Error("iteration over array fail")
		}

		if t1.K("num").ForEachArr(func(o Obj, idx int) {
			t.Error("iteration over non-array fail")
		}) == nil {
			t.Error("iteration over non-array fail")
		}

	}
	{ //ForEachObj
		if t1.K("null").ForEachObj(func(o Obj, key string) {
			t.Error("iteration over null fail")
		}) != nil {
			t.Error("iteration over null fail")
		}

		count := 0
		if t1.K("obj").ForEachObj(func(o Obj, key string) {
			count++
			if key != "a" && key != "b" {
				t.Error("object iteration key fail")
			}
			if (key == "a" && o.Int(nil) != 1) ||
				(key == "b" && o.Int(nil) != 2) {
				t.Error("object iteration fail")
			}
		}) != nil {
			t.Error("iteration over array fail")
		}
		if count != 2 {
			t.Error("iteration over array fail")
		}

		if t1.K("num").ForEachObj(func(o Obj, key string) {
			t.Error("iteration over non-object fail")
		}) == nil {
			t.Error("iteration over non-object fail")
		}

	}
}

func TestModifiers(t *testing.T) {
	t1 := FromJSON(`{"a":[]}`)

	{ //Push
		t2 := t1.K("a")
		t2.Push("zz")
		fmt.Println(t1.String(nil))
		if t2.Err != nil || t2.Len() != 1 || t1.K("a").I(0).String(nil) != "zz" {
			fmt.Println("t1=", t1.String(nil))
			fmt.Println("t2=", t2.String(nil))
			t.Error("Push error")
		}

		//non array
		t1.Push("zz")
		if t1.Err == nil || t1.K("a").IsNil() {
			t.Error("Push error")
		}
	}

}
