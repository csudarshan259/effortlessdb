package main

import (
	"effortlessdb/others"
	"effortlessdb/others/create"
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args

	if len(arguments) == 0 {
		os.Exit(0)
	}

	identifier := strings.ToLower(arguments[1])
	//fmt.Println(identifier)
	if identifier == others.Create {

		isSuccess, result := create.AddKeyValueDoc(arguments[2], arguments[3], arguments[4])
		fmt.Println(result, isSuccess)

	} else if identifier == others.Read {

	} else if identifier == others.Update {

	} else if identifier == others.Delete {

	} else {
		os.Exit(0)
	}

}
