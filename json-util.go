//Package jsonutil : wrapper to help handling JSON objects
package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
	"strings"
)

//RawObj Type alias for the Go representation of a JSON object
type RawObj = map[string]interface{}

//RawArray Type alias for the Go representation of a JSON array
type RawArray = []interface{}

//Obj : the basic JSON wrapper object
type Obj struct {
	// root interface{}
	Data interface{}
	Err  error
}

//DBG variable to control the level of debug logging
var DBG struct {
	PrintDetail  bool
	ShowWarnings bool
}

//Load create the chainable object from
func Load(obj interface{}) Obj {
	if obj == nil {
		if DBG.ShowWarnings {
			log.Println("JSONUtil: nil object received")
		}
		return Obj{Err: errors.New("JSONUtil Nil object received")}
	}
	return Obj{
		Data: obj,
		Err:  nil,
	}
}

//FromJSON creates a JSON representation from a byte slice
func FromJSON(jsonInp []byte) (out Obj) {
	out = Obj{}

	if err := json.Unmarshal(jsonInp, &out.Data); err != nil {
		out.Err = err
	}
	return
}

//IsNil : checks if the raw data is a `nil`
func (o Obj) IsNil() bool { return o.Data == nil }

//Raw : this is for testing purposes only, returns the raw data object
func (o Obj) Raw() interface{} { return o.Data }

//K : walk one, or more steps deeper into the structure  K=key
func (o Obj) K(keys ...string) (n Obj) {
	defer func() {
		if r := recover(); r != nil {
			n = Obj{Err: r.(error)}
			if DBG.ShowWarnings {
				log.Println("JSON walk error", r)
			}
			if DBG.PrintDetail {
				debug.PrintStack()
			}
		}
	}()

	n = Obj{Data: o.Data}

	for _, idx := range keys {
		n.Data = n.Data.(RawObj)[idx]
	}
	return n
}

//I : walk one, or more dimension down into the array
func (o Obj) I(index ...int) (n Obj) {
	defer func() {
		if r := recover(); r != nil {
			n = Obj{Err: r.(error)}
			if DBG.ShowWarnings {
				log.Println("JSON walk error", r)
			}
			if DBG.PrintDetail {
				debug.PrintStack()
			}
		}
	}()

	n = Obj{Data: o.Data}

	for _, idx := range index {
		n.Data = n.Data.(RawArray)[idx]
	}
	return n
}

//String : convert value to a string
func (o Obj) String() (s string) {
	defer func() {
		if r := recover(); r != nil {
			s = ""
			if DBG.ShowWarnings {
				log.Println("JSON to string conversion error:", r)
			}
			if DBG.PrintDetail {
				debug.PrintStack()
			}
		}
	}()

	s = ""
	if o.Data == nil {
		return s
	}
	switch v := o.Data.(type) {
	case float64:
		s = strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		s = v
	case bool:
		if v {
			s = "true"
		} else {
			s = "false"
		}
	default:
		s = fmt.Sprintf("%+v", o.Data) //generate a full printout
	}
	return s
}

//Int64 converts value to an Int64
func (o Obj) Int64() (i int64) {
	defer func() {
		if r := recover(); r != nil {
			i = 0
			if DBG.ShowWarnings {
				log.Println("JSON to Int64 conversion error:", r)
			}
			if DBG.PrintDetail {
				debug.PrintStack()
			}

		}
	}()
	o.Err = nil
	i = 0
	switch v := o.Data.(type) {
	case float64:
		i = int64(v)
	case int:
		i = int64(v)
	case float32:
		i = int64(v)
	case uint:
		i = int64(v)
	case int32:
		i = int64(v)
	case uint32:
		i = int64(v)
	case int64:
		i = int64(v)
	case uint64:
		i = int64(v)
	case string:
		x, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			o.Err = errors.New(err.Error())
		}
		i = int64(x)
	default:
		i = o.Data.(int64) //generate an error
	}
	return i
}

//Int converts value to an Int
func (o Obj) Int() (i int) {
	return int(o.Int64())
}

//Float64 converts value to an Float64
func (o Obj) Float64() (f float64) {
	defer func() {
		if r := recover(); r != nil {
			f = 0
			if DBG.ShowWarnings {
				log.Println("JSON to Float64 conversion error:", r)
			}
			if DBG.PrintDetail {
				debug.PrintStack()
			}

		}
	}()
	o.Err = nil
	f = 0
	switch v := o.Data.(type) {
	case float64:
		f = float64(v)
	case int:
		f = float64(v)
	case float32:
		f = float64(v)
	case uint:
		f = float64(v)
	case int32:
		f = float64(v)
	case uint32:
		f = float64(v)
	case int64:
		f = float64(v)
	case uint64:
		f = float64(v)
	case string:
		x, err := strconv.ParseFloat(v, 64)
		if err != nil {
			o.Err = errors.New(err.Error())
		}
		f = float64(x)
	default:
		f = o.Data.(float64) //generate an error
	}
	return f
}

//Float32 converts value to an Float32
func (o Obj) Float32() float32 {
	return float32(o.Float64())
}

//Bool converts value to an Bool
func (o Obj) Bool() (b bool) {
	defer func() {
		if r := recover(); r != nil {
			b = false
			if DBG.ShowWarnings {
				log.Println("JSON to Bool conversion error:", r)
			}
			if DBG.PrintDetail {
				debug.PrintStack()
			}
		}
	}()
	b = false
	if o.Data == nil {
		return b
	}

	switch v := o.Data.(type) {
	case bool:
		b = v
	case string:
		v = strings.ToLower(v)
		if v == "true" || v == "yes" {
			return true
		}
		if v == "false" || v == "no" {
			return false
		}
		b = o.Data.(bool) //generate an error
	default:
		b = o.Data.(bool) //generate an error
	}
	return b
}

//ForEachArr ITERATE OVER AN ARRAY
func (o Obj) ForEachArr(F func(Obj, int)) {
	defer func() {
		if r := recover(); r != nil {
			if DBG.ShowWarnings {
				log.Println("ARRAY iteration error:", r)
			}
			if DBG.PrintDetail {
				debug.PrintStack()
			}
		}
	}()

	if o.Data == nil {
		return
	}
	it := o.Data.(RawArray) //it will throw an error automatically if it's not the correct type
	for idx, val := range it {
		F(Obj{Data: val}, idx)
	}
}

//Push adds an object to an array
func (o *Obj) Push(value interface{}) {
	defer func() {
		if r := recover(); r != nil {
			if DBG.ShowWarnings {
				log.Println("Not an array was used in Push operation:", r)
			}
		}
	}()

	o.Data = append(o.Data.(RawArray), value)
}

//ForEachObj ITERATE OVER AN OBJECT
func (o Obj) ForEachObj(F func(Obj, string)) {
	defer func() {
		if r := recover(); r != nil {
			if DBG.ShowWarnings {
				log.Println("OBJECT iteration error:", r)
			}
			if DBG.PrintDetail {
				debug.PrintStack()
			}
		}
	}()

	if o.Data == nil {
		return
	}
	it := o.Data.(RawObj) //it will throw an error automatically if it's not the correct type
	for idx, val := range it {
		F(Obj{Data: val}, idx)
	}
}

// Len : The length (size of an array or object)
// if the object is no an array or object then result is -1
func (o Obj) Len() int {
	switch v := o.Data.(type) {
	case RawArray:
		return len(v)
	case RawObj:
		return len(v)
	}
	return -1
}

//ToJSON GENERATES THE JSON STRING REPRESENTATION OF THE OBJECT
func (o Obj) ToJSON() string {

	out, err := json.Marshal(o.Data)
	if err != nil {
		return "ERROR"
	}
	return string(out)
}

// AddKV adds a key/value pair to the JSON object
func (o Obj) AddKV(key string, value interface{}) (n Obj) {
	defer func() {
		if r := recover(); r != nil {
			if DBG.ShowWarnings {
				log.Println("OBJECT iteration error:", r)
			}
			n = Obj{Err: r.(error)}
			if DBG.PrintDetail {
				debug.PrintStack()
			}
		}
	}()

	o.Data.(RawObj)[key] = value
	return o
}

// RemoveKey removes a key from a JSON object
func (o Obj) RemoveKey(key string) (n Obj) {
	defer func() {
		if r := recover(); r != nil {
			if DBG.ShowWarnings {
				log.Println("RemoveKey error:", r)
			}
			n = Obj{Err: r.(error)}
			if DBG.PrintDetail {
				debug.PrintStack()
			}
		}
	}()

	delete(o.Data.(RawObj), key)

	return o
}

//ToStringArray converts the array in the object to a []string type
func (o Obj) ToStringArray() (out []string) {
	defer func() {
		if r := recover(); r != nil {
			if DBG.ShowWarnings {
				log.Println("RemoveKey error:", r)
			}
			out = []string{}
			o.Err = r.(error)
			if DBG.PrintDetail {
				debug.PrintStack()
			}
		}
	}()

	x1 := o.Data.([]interface{})

	out = make([]string, len(x1))
	for i := range x1 {
		out[i] = x1[i].(string)
	}
	return out
}
