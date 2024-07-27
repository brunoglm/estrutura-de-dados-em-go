package main

import "fmt"

func main() {
	hashMap := map[string]string{
		"key1": "value1",
	}

	hashMap["key2"] = "value2"

	hashMap["key1"] = "valorAlterado"

	fmt.Println(hashMap["key1"])

	delete(hashMap, "key1")

	fmt.Println(hashMap)
}
