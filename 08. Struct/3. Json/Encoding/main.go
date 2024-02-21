package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool // #3

type user struct { // #1
	//Json package only encodes exported fieds nên cần viết hoa
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"` // #6	fied tags are also part of a struct's type. A field tech has a key and a value.
	// Key Name: Json. The json package will read this tag
	// Value: perms, omitempty. perms: change the encoded field name, omits the field if it has
	// name        string // #1
	// password    string // #1
	// permissions // #3
}

func main() {
	users := []user{ // #2
		{"inanc", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "66", permissions{"write": true}},
	}

	// out, err := json.Marshal(users) // #4
	out, err := json.MarshalIndent(users, "", "\t") // #5 Marshal return bytes. MarshalIndent sử dụng để in ra cho đẹp dễ đọc. Bthg sử dụng Marshal
	if err != nil {                                 // #4
		fmt.Println(err)
		return
	}

	fmt.Println(string(out)) // #4
}
