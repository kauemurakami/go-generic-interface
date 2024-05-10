[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-generic-interface/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-generic-interface/blob/main/README.md)  
go version 1.22.1  

## Generic Interfaces 
Using the interface as a generic type, be very careful not to turn it into a hack.
Unlike common interfaces, a generic interface is served by anything, without restrictions, let's see:  
```go
package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("Kauê") // output Kauê
	generic(1) // output 1
	generic(true) // output true
}
```
We created a function with an interface type parameter, this means that it can receive a parameter that meets the passed interface, but as there is nothing in the interface, EVERYTHING meets the interface, a string, an int, etc. And as we saw, we can basically pass everything on to her.  
Additionally, it can be used with a ```map``` for example:  
```go
func main(){
...
	map0 := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}
	fmt.Println(map0) //output map[String:String 1:String 100:true true:12]
}
```
But be careful not to create a huge problem in your code.