# go-json-util

Simplistic, crash free handling of JSON in Golang.

I have created this module in order to deal with JSON based API calls, which I frequently handle on the projects I'm working with. 

I was very tired of all the boilerplating, and typecasting and runtime crashes when the typecasting failed. So I looked around for a simple module I could use. But found none that met my criteria:

- Need very simple mechanism for converting to and from JSON strings/byte stream
- Concise and simple syntax to traverse the JSON objects and get data out of them (or to modify them)
- Should not ever crash, even if keys don't exist, or type of data does not match
- Should do type conversion automatically, if possible when accessing data 

So, I build my solution based on these principle and have been using it in production for over a year, and it made my life quite a bit easier. So, I'm sharing it now for others who want simple reliable handling for JSON.

The module does not contain all JSON related operations. It only has what I needed on my projects but I think it covers probably 99% of cases.

## Basic concept 

`json-util` is utilizing a simple wrapper around the raw JSON data turned into type `interface{}` (in other words, Go's generic type), called `Obj`

The methods on this type allow you to:
- Extract a value for basic types (numeric, string, bool), and check for null/nil
- Traverse the structure in a chainable manner (see later for examples)
- Add or remove keys in objects/maps
- Iterate over arrays and objects (maps)
- Convert it to a JSON string

You can create the `Obj` from an existing variable of any type (a struct, or map, etc) or by parsing in a JSON string.

You can also get the raw Go data out. I found it necessary in some cases when I needed to pass it to another library's method.

### Error handling 

Methods never fail (at least not supposed to) and always return a value when you try retrieve data. When error occurs (eg: accessing a non-existent key) the returned value is the usual initial value for the type (eg: 0 for numbers, "" for strings, `false` for booleans). 

If you want to check for errors then simply check the `.Err` attribute of the `Obj` after the operation. It's of `error` type. 

### Debugging

This module is designed to be silent and don't complain no matter what. 

But sometimes it's necessary to see why thing go haywire, and where do they go haywire. 

To can configure the verbosity by setting the `DBG` variable in the module:

`jsonutil.DBG.ShowWarnings=true` - This will show warning messages when they happen.

`jsonutil.DBG.PrintDetail=true` - This will print the stack trace when errors are catched. Useful for tracking down the offending part of your code.


## Examples





# Reference


### type Obj

```go
type Obj struct {
	Data interface{}
	Err  error
}
```

Obj : the basic JSON wrapper object

### func  FromJSON

```go
func FromJSON(jsonInp string) (out Obj)
```
FromJSON creates a JSON representation from a byte slice

### func  Load

```go
func Load(obj interface{}) Obj
```
Load creates the chainable object from the variable passed to it

### func (*Obj) AddKV

```go
func (o *Obj) AddKV(key string, value interface{})
```
AddKV adds a key/value pair to the JSON object

### func (Obj) Bool

```go
func (o Obj) Bool(toErr *error) (b bool)
```
Bool converts value to a Bool

### func (Obj) Float32

```go
func (o Obj) Float32(toErr *error) float32
```
Float32 converts value to an Float32

### func (Obj) Float64

```go
func (o Obj) Float64(toErr *error) (f float64)
```
Float64 converts value to an Float64

### func (Obj) ForEachArr

```go
func (o Obj) ForEachArr(F func(Obj, int)) (toErr error)
```
ForEachArr ITERATE OVER AN ARRAY

### func (Obj) ForEachObj

```go
func (o Obj) ForEachObj(F func(Obj, string)) (toErr error)
```
ForEachObj ITERATE OVER AN OBJECT

### func (Obj) I

```go
func (o Obj) I(index ...int) (n Obj)
```
I : walk one, or more dimension down into an array

### func (Obj) Int

```go
func (o Obj) Int(toErr *error) (i int)
```
Int converts value to an Int

### func (Obj) Int64

```go
func (o Obj) Int64(toErr *error) (i int64)
```
Int64 converts value to an Int64

### func (Obj) IsNil

```go
func (o Obj) IsNil() bool
```
IsNil : checks if the raw data is a `nil`

### func (Obj) K

```go
func (o Obj) K(keys ...string) (n Obj)
```
K : walk one, or more steps deeper into the structure K=key

### func (Obj) Len

```go
func (o Obj) Len() int
```
Len : The length (size of an array or object) if the object is no an array or
object then result is -1

### func (*Obj) Push

```go
func (o *Obj) Push(value interface{})
```
Push adds an object to an array

### func (Obj) Raw

```go
func (o Obj) Raw() interface{}
```
Raw : this is for testing purposes only, returns the raw data object

### func (Obj) RemoveKey

```go
func (o Obj) RemoveKey(key string) (n Obj)
```
RemoveKey removes a key from a JSON object

### func (Obj) String

```go
func (o Obj) String(toErr *error) (s string)
```
String : convert value to a string

### func (Obj) ToJSON

```go
func (o Obj) ToJSON() string
```
ToJSON GENERATES THE JSON STRING REPRESENTATION OF THE OBJECT

### func (Obj) ToStringArray

```go
func (o Obj) ToStringArray(toErr *error) (out []string)
```
ToStringArray converts the array in the object to a []string type

### type RawArray

```go
type RawArray = []interface{}
```

RawArray Type alias for the Go representation of a JSON array

### type RawObj

```go
type RawObj = map[string]interface{}
```

RawObj Type alias for the Go representation of a JSON object
