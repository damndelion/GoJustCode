package main

import (
	"encoding/json"
	"fmt"
)

func printJSONStructure(data interface{}, prefix string) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			fmt.Printf("%s%s (object)\n", prefix, key)
			printJSONStructure(value, prefix+"  ")
		}
	case []interface{}:
		for i, item := range v {
			fmt.Printf("%s[%d] (array)\n", prefix, i)
			printJSONStructure(item, prefix+"  ")
		}
	default:
		fmt.Printf("%s(%T)\n", prefix, v)
	}
}

func main() {
	// Замените эту строку на ваш JSON-код
	jsonData := `{"name": "John", "age": 30, "city": "New York", "pets": ["dog", "cat"]}`

	var data interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	fmt.Println("Структура JSON:")
	printJSONStructure(data, "")

}
