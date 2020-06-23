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

//Obj is the basic JSON wrapper object
type Obj struct {
	Data    interface{}
	Err     error
	updater func(interface{}) //function with closure to update the value without breaking references
}

//DBG variable to control the level of debug logging
var DBG struct {
	PrintDetail  bool
	ShowWarnings bool
}

/*████████████████████████████████████████████
██████ HELPER FUNCTIONS
████████████████████████████████████████████*/

//recovery Internal function to run the repeating code for handling cases when things would panic
func recovery(warn string, setter func(error)) {
	if r := recover(); r != nil {
		setter(r.(error))
		if DBG.ShowWarnings {
			log.Println(warn, r)
		}
		if DBG.PrintDetail {
			debug.PrintStack()
		}
	}
}

func setError(to *error, newErr error) {
	if to != nil {
		*to = newErr
	}
}

/*████████████████████████████████████████████
██████ JSON CONVERSIONS + CONSTRUCTORS
████████████████████████████████████████████*/

//Load creates the chainable object from the variable passed to it
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
func FromJSON(jsonInp string) (out Obj) {
	out = Obj{}

	if err := json.Unmarshal([]byte(jsonInp), &out.Data); err != nil {
		out.Err = err
	}
	return
}

//ToJSON generates the json string representation of the object
func (o Obj) ToJSON() string {

	out, err := json.Marshal(o.Data)
	if err != nil {
		return "ERROR"
	}
	return string(out)
}

/*████████████████████████████████████████████
██████ WALK METHODS, to access nested, indexed content
████████████████████████████████████████████*/

//K  walk one, or more steps deeper into the structure  K=key
func (o Obj) K(keys ...string) (n Obj) {
	defer recovery("JSON walk error", func(err error) {
		n = Obj{Err: err}
	})

	o.Err = nil
	n = Obj{Data: o.Data}

	for _, idx := range keys {
		encaps := n.Data.(RawObj)
		n.updater = func(val interface{}) {
			encaps[idx] = val
		}
		n.Data = n.Data.(RawObj)[idx]
	}
	return n
}

//I = walk one, or more dimension down into an array (I = index)
func (o Obj) I(index ...int) (n Obj) {
	defer recovery("JSON walk error", func(err error) {
		n = Obj{Err: err}
	})

	o.Err = nil
	n = Obj{Data: o.Data}

	for _, idx := range index {
		n.updater = func(val interface{}) {
			n.Data.(RawArray)[idx] = val
		}
		n.Data = n.Data.(RawArray)[idx]
	}
	return n
}

/*████████████████████████████████████████████
██████ CONVERSION METHODS
████████████████████████████████████████████*/

//String convert value to a string
func (o Obj) String(toErr *error) (s string) {
	setError(toErr, nil)
	defer recovery("JSON to string conversion error:", func(err error) {
		setError(toErr, err)
	})

	if o.Data == nil {
		setError(toErr, errors.New(" `nil` to string conversion"))
		return ""
	}
	return fmt.Sprintf("%+v", o.Data) //use fmt's conversion for pretty much any type
}

//Int64 converts value to an Int64
func (o Obj) Int64(toErr *error) (i int64) {
	setError(toErr, nil)
	defer recovery("JSON to int conversion error:", func(err error) {
		setError(toErr, err)
	})

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
			setError(toErr, err)
		}
		i = int64(x)
	default:
		i = o.Data.(int64) //generate an error
	}
	return i
}

//Int converts value to an Int
func (o Obj) Int(toErr *error) (i int) {
	return int(o.Int64(toErr))
}

//Float64 converts value to an Float64
func (o Obj) Float64(toErr *error) (f float64) {
	setError(toErr, nil)
	defer recovery("JSON to Float64 conversion error:", func(err error) {
		setError(toErr, err)
	})

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
			setError(toErr, err)
		}
		f = float64(x)
	default:
		f = o.Data.(float64) //generate an error
	}
	return f
}

//Float32 converts value to an Float32
func (o Obj) Float32(toErr *error) float32 {
	return float32(o.Float64(toErr))
}

//Bool converts value to a Bool
func (o Obj) Bool(toErr *error) (b bool) {
	setError(toErr, nil)
	defer recovery("JSON to Bool conversion error:", func(err error) {
		setError(toErr, err)
	})

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
		setError(toErr, errors.New("can't convert value to Bool"))
	default:
		setError(toErr, errors.New("can't convert value to Bool"))
	}
	return b
}

//ToStringArray converts the array in the object to a []string type
func (o Obj) ToStringArray(toErr *error) (out []string) {
	setError(toErr, nil)
	defer recovery("ToStringArray error:", func(err error) {
		setError(toErr, err)
	})

	out = []string{}
	out = make([]string, len(o.Data.([]interface{})))
	i := 0
	*toErr = o.ForEachArr(func(val Obj, _ int) {
		out[i] = val.String(nil)
		i++
	})

	return out
}

/*████████████████████████████████████████████
██████ ITERATORS
████████████████████████████████████████████*/

//ForEachArr iterate over an array, calling a function with the values
func (o Obj) ForEachArr(F func(Obj, int)) (toErr error) {
	defer recovery("ARRAY iteration error:", func(err error) {
		toErr = err
	})

	if o.Data == nil {
		return nil
	}
	it := o.Data.(RawArray) //it will throw an error automatically if it's not the correct type
	for idx, val := range it {
		F(Obj{Data: val,
			updater: func(newVal interface{}) {
				it[idx] = newVal
			}}, idx)
	}
	return nil
}

//ForEachObj iterate over an object, calling a function with the values
func (o Obj) ForEachObj(F func(Obj, string)) (toErr error) {
	defer recovery("OBJECT iteration error:", func(err error) {
		toErr = err
	})

	if o.Data == nil {
		return nil
	}
	it := o.Data.(RawObj) //it will throw an error automatically if it's not the correct type
	for idx, val := range it {
		F(Obj{Data: val,
			updater: func(newVal interface{}) {
				it[idx] = newVal
			}}, idx)
	}
	return nil
}

/*████████████████████████████████████████████
██████ MODIFIER METHODS
████████████████████████████████████████████*/

//Push adds an object to an array
func (o *Obj) Push(value interface{}) {
	defer recovery("Push error:", func(err error) {
		o.Err = err
	})

	o.Err = nil
	o.Data = append(o.Data.(RawArray), value)
	if o.updater != nil {
		o.updater(o.Data)
	}
}

// AddKV adds a key/value pair to the JSON object
func (o *Obj) AddKV(key string, value interface{}) {
	defer recovery("AddKV error:", func(err error) {
		o.Err = err
	})

	o.Err = nil

	o.Data.(RawObj)[key] = value
}

// RemoveKey removes a key from a JSON object
func (o Obj) RemoveKey(key string) (n Obj) {
	defer recovery("RemoveKey error:", func(err error) {
		n = Obj{Err: err}
	})

	o.Err = nil

	delete(o.Data.(RawObj), key)

	return o
}

/*████████████████████████████████████████████
██████ INFO METHODS
████████████████████████████████████████████*/

// Len gives the length (size of an array or object)
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

//IsNil checks if the raw data is a `nil`
func (o Obj) IsNil() bool { return o.Data == nil }

//Raw returns the raw Golang interface object for the JSON object
func (o Obj) Raw() interface{} { return o.Data }
