package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("Kauê")
	generic(1)
	generic(true)

	map0 := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}
	fmt.Println(map0)

}
