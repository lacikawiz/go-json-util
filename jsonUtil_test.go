package jsonutil

import (
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Run("Array with objects", func(t *testing.T) {
		t1 := FromJSON([]byte(`[{"x":1},{"z":"str"},{"y":true}]`))
		if t1.I(0).K("x").Int() != 1 {
			t.Error("int test failed")
		}
		if t1.I(1).K("z").String() != "str" {
			t.Error("string test failed")
		}
		if !t1.I(2).K("y").Bool() {
			t.Error("bool test failed")
		}
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		t2 := FromJSON([]byte(`{x:1}`))
		if t2.Err == nil {
			t.Error("invalid JSON test failed")
		}
		t2 = FromJSON([]byte(`[w]`))
		if t2.Err == nil {
			t.Error("invalid JSON test failed")
		}
	})

	t3 := Obj{}
	t.Run("Nested Objects", func(t *testing.T) {
		t3 = FromJSON([]byte(`{"x":{
			"a":2,"b":"str","c":[1,2,"3"]
		}}`))
		if t3.K("x", "c").I(2).String() != "3" {
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
	t1 := FromJSON([]byte(testStr))
	res := t1.ToJSON()
	if res != testStr {
		t.Error("JSON encoding test failed:")
		t.Log("Should have:", testStr)
		t.Log("Got", res)
	}
}

func TestQueryString(t *testing.T) {
	t1 := FromJSON([]byte(`{"x#":{
		"a":2,"b":"str","c!":[1,2,"3"]
	}}`))
	res := t1.ToQueryString()
	//the sequence is random so we need to test it this way
	if !(strings.Contains(res, "x%23[c%21][]=1") &&
		strings.Contains(res, "x%23[c%21][]=2") &&
		strings.Contains(res, "x%23[c%21][]=3") &&
		strings.Contains(res, "x%23[a]=2") &&
		strings.Contains(res, "x%23[b]=str") &&
		len(res) == 66) {
		t.Error("Query building test failed")
	}

	t2 := Load(map[string]interface{}{
		"filter": map[string]interface{}{"EMAIL": "laszlo@laszlo.com"},
		"select": []interface{}{"EMAIL", "*", "PHONE"},
	})
	res = t2.ToQueryString()
	if !(strings.Contains(res, "filter[EMAIL]=laszlo@laszlo.com") &&
		strings.Contains(res, "select[]=EMAIL") &&
		strings.Contains(res, "select[]=%2A") &&
		strings.Contains(res, "select[]=PHONE") &&
		len(res) == 74) {
		t.Error("Query building from literal failed")
		t.Log(res)
	}

}

func TestAddKV(t *testing.T) {
	t1 := FromJSON([]byte(`[{"x":1},{"z":"str"},{"y":true}]`))
	t1.I(0).AddKV("long", "wait").AddKV("short", -10)

	if t1.I(0).K("long").String() != "wait" ||
		t1.I(0).K("short").Int() != -10 {
		t.Error("Adding Key Value failed")
	}
}