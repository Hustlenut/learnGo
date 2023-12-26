package misc

import (
	"encoding/json"
	"fmt"
)

/*
	Struct tags are annotations added to struct fields
	to provide metadata. They help processes like JSON marshaling
	(encoding) and unmarshaling (decoding)
	interpret and handle struct fields correctly. For example, in
	JSON encoding, struct tags can specify the field names in
	the resulting JSON object. In this case the field "F1" is a
	string, that will be represented as the key "field1 in json.
	The "omitempty" struct tag option in Go is used to control how
	struct fields are marshaled (encoded) into JSON when using the
	encoding/json package. When a field is tagged with omitempty,
	it instructs the JSON encoder to omit (exclude) the field from
	the JSON output if the field has a zero or empty value in Go.
*/
type T struct {
	F1 string `json:"field1"`
	F2 string `json:"field2",omitempty`
	F3 string `json:"field3",omitempty`
	F4 string `json:"-"`
}

func PrintStructTags() {
	fmt.Println("------------")
	fmt.Println("Struct Tags")

	t := T{
		F1: "v1",
		F2: "",
		F3: "v3",
		F4: "v4",
	}
	/*
		The json.Marshal function returns two values:
		The first return value ("result" in code) is the JSON
		representation of the data structure (t) as a byte
		slice ([]byte).
		The second return value is an error type.
		It indicates whether an error occurred during the
		marshaling process. If there was no error, this value
		will be nil. However, if an error occurred
		(e.g., due to an invalid data structure or encoding issue),
		it will contain an error message.
	*/
	result, _ := json.Marshal(t)

	fmt.Printf("%s\n", result) //"\n" leaves a line break at the end of this message.
	fmt.Println("------------")
}
