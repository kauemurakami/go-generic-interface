[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-generic-interface/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-generic-interface/blob/main/README.md)  
go version 1.22.1  

## Interfaces Genéricas 
Usar a interface como sendo um tipo genérico, tomar muito cuidado pra não virar gambiarra.  
Diferente das interfaces comuns, uma interface genérica é atendida por qualquer coisa, sem restrições, vejamos:  
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
Criamos uma funçõa com um parametro do tipo interface, isso quer dizer que ela pode receber um parâmetro que atenda a interface passado, mas como não tem nada na interface, TUDO atendo a interface, uma string, um int e etc. E como vimos, basicamente podemos passar tudo para ela.  
Além disso pode ser usada com um  ```map``` por exemplo:  
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
Mas cuidado pra isso não gerar uma grande gambiarra no seu código.  