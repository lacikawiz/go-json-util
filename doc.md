
---

#### module variables

```go
var DBG struct {
	PrintDetail  bool
	ShowWarnings bool
}
```
DBG variable to control the level of debug logging

---

#### type Obj

```go
type Obj struct {
	Data interface{}
	Err  error
}
```

Obj is the basic JSON wrapper object

---

#### func  FromJSON

```go
func FromJSON(jsonInp string) (out Obj)
```
FromJSON creates a JSON representation from a byte slice

---

#### func  Load

```go
func Load(obj interface{}) Obj
```
Load creates the chainable object from the variable passed to it

---


#### func (Obj) Bool

```go
func (o Obj) Bool(toErr *error) (b bool)
```
Bool converts value to a Bool

---

#### func (Obj) Float32

```go
func (o Obj) Float32(toErr *error) float32
```
Float32 converts value to an Float32

---

#### func (Obj) Float64

```go
func (o Obj) Float64(toErr *error) (f float64)
```
Float64 converts value to an Float64

---

#### func (Obj) ForEachArr

```go
func (o Obj) ForEachArr(F func(Obj, int)) (toErr error)
```
ForEachArr iterate over an array, calling a function with the values

---

#### func (Obj) ForEachObj

```go
func (o Obj) ForEachObj(F func(Obj, string)) (toErr error)
```
ForEachObj iterate over an object, calling a function with the values

---

#### func (Obj) I

```go
func (o Obj) I(index ...int) (n Obj)
```
I = walk one, or more dimension down into an array (I = index)

---

#### func (Obj) Int

```go
func (o Obj) Int(toErr *error) (i int)
```
Int converts value to an Int

---

#### func (Obj) Int64

```go
func (o Obj) Int64(toErr *error) (i int64)
```
Int64 converts value to an Int64

---

#### func (Obj) IsNil

```go
func (o Obj) IsNil() bool
```
IsNil checks if the raw data is a `nil`

---

#### func (Obj) K

```go
func (o Obj) K(keys ...string) (n Obj)
```
K walk one, or more steps deeper into the structure K=key

---

#### func (Obj) Len

```go
func (o Obj) Len() int
```
Len gives the length (size of an array or object) if the object is no an array
or object then result is -1

---

#### func (*Obj) Push

```go
func (o *Obj) Push(value interface{})
```
Push adds an object to an array

---

#### func (Obj) Raw

```go
func (o Obj) Raw() interface{}
```
Raw returns the raw Golang interface object for the JSON object

---

#### func (Obj) RemoveKey

```go
func (o Obj) RemoveKey(key string) (n Obj)
```
RemoveKey removes a key from a JSON object

---

#### func (Obj) String

```go
func (o Obj) String(toErr *error) (s string)
```
String convert value to a string

---

#### func (Obj) ToJSON

```go
func (o Obj) ToJSON() string
```
ToJSON generates the json string representation of the object

---

#### func (Obj) ToStringArray

```go
func (o Obj) ToStringArray(toErr *error) (out []string)
```
ToStringArray converts the array in the object to a []string type

---

#### type RawArray

```go
type RawArray = []interface{}
```

RawArray Type alias for the Go representation of a JSON array

---

#### type RawObj

```go
type RawObj = map[string]interface{}
```

RawObj Type alias for the Go representation of a JSON object

---

#### func (*Obj) AddKV

```go
func (o *Obj) AddKV(key string, value interface{})
```
AddKV adds a key/value pair to the JSON object

