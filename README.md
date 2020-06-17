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

## json-util.AddKV
  AddKV adds a key/value pair to a map. If the object is not a map then prints an error and doesn't do anything. Existing keys are overwritten.

  `func (o Obj) AddKV(key string,value interface{}) (n Obj)`

## json-util.RemoveKey
  Removes a key from a map.
  `func (o Obj) RemoveKey(key string) (n Obj)`

## json-util.ForEachArr
  Iterate over an array.

  Calls the given function for each ''value'', and ''index'' of the array. If the object is not an array, then it simply doesn't do anything.

  `func (o Obj) ForEachArr(F func(Obj,int))`

## json-util.ForEachObj
  Iterate over a map in the JSON object.

  Calls the given function for each ''value'', and ''key'' of the map. If the object is not a map, then it simply doesn't do anything.

  `func (o Obj) ForEachObj(F func(Obj,string))`

## json-util.FromJSON
  Converts the JSON encoded input into the json-util object. If there's any error it can be retrieved from the `.Err` of the output variable.

  `func FromJSON(jsonInp []byte) (out Obj)`

## json-util.ToJSON
  Returns the JSON encoded value of the object.

  `func (o Obj) ToJSON() string`

## json-util.I
  Walk one, or more dimension down into the array

  `func (o Obj) I(index ...int) (n Obj)`

  > Example:
  ``val:=J.FromJSON(`[1,[2,3]`)
  cVal:=val.I(1,1).Int()   //should be 3
  ``

## json-util.IsNil
  Returns true is the json-util object contains a nil value

  `func (o Obj) IsNil() bool`

## json-util.Len
  Returns the length (size of an array or map) if the object is no an array or map then result is -1

  `func (o Obj) Len() int`

## json-util.Load
  Adds raw Golang data a json-util object, so that further functions can work with it.

  `func Load(obj interface{}) Obj`

  The input can be any type but for proper operation, limit them to the following types:

  * Basic types
  * map[string]interface{}
  * []interface{}

## json-util.Raw
  return the raw data object

  `func (o Obj) Raw() interface{} {return o.Data }`

## json-util.Obj (Type)
  The type which is used by this module

  `type Obj struct {
    Data interface{}
    Err error
  }`

  > `Data`: Contains the raw data object
  ---
  > `Err`: The error that happened during the last action (used this rather than each function returning an error too)

## json-util.Push
  Adds a value to an Array. If the object is not an array then it prints an error message and doesn't do anything.

  `func (o *Obj) Push(value interface{})`

## Conversion Functions

  These are the extraction/conversion functions. They all attempt to make all possible effort to convert the value to the requested format (eg: int to string, string to int, etc). If not successful then they print a debug log info and return the default value for the type.

  * `func (o Obj) String() (s string)`
  * `func (o Obj) Int64() (i int64)`
  * `func (o Obj) Int() (i int)`
  * `func (o Obj) Float64() (f float64)`
  * `func (o Obj) Float32() float32`
  * `func (o Obj) Bool() (b bool)`

