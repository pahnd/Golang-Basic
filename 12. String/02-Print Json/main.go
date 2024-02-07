package main

import "fmt"

func main() {
	// this one uses a raw string literal

	// can you see how readable it is?
	// compared to the previous one?

	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"

	json1 := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`
	fmt.Println(json)
	fmt.Println(json1)
}
