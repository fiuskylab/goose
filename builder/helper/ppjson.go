package helper

import (
	"encoding/json"
	"fmt"
)

func PPJson(v any) {
	b, _ := json.MarshalIndent(v, "", "			")
	fmt.Println(string(b))
}
