package main

import (
	"fmt"
	"encoding/json"
	"syscall/js"
)

func preetyJSON(input string) (string, error) {
	var raw interface{}
	err := json.Unmarshal([]byte(input), &raw)
	if err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "   ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments"
		}

		jsonString := args[0].String()
		fmt.Printf("input: %s\n", jsonString)

		pretty, err := preetyJSON(jsonString)

		if err != nil {
			fmt.Printf("Unable to convert to JSON %s\n", err)
			return err.Error()
		}
		return pretty

	})

	return jsonFunc
}

func main() {
	fmt.Println(jsonWrapper())
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan bool)
}
