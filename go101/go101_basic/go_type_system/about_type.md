### 1、byte is a built-in alias of uint8, rune is a built-in alias of int32
### 2、type NewType SourceType
- a: NewType and SourceType are two distinct types // 两种不同的类型
- b: NewType and SourceType can be converted to each other
- c: type can be defined within function body
### 3、type Name = string
- Name has the same type of string
### 4、nil is the default value of slice,map,function,channel,pointer and interface
### 5、signature of function type
- the signature of a function type is composed of the input parameter list and the output result list, not 
include the function name and body
### support or not support comparison
#### the following types are not comparison
- slice type
- map type
- function type
- any struct type with a field above
- **the other types are all comparison**

## pointer
- return pointer of local variables is safe in go