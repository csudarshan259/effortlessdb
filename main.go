package main

import (
	"effortlessdb/Others"
	"effortlessdb/Others/create"
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
	if identifier == Others.Create {

		isSuccess, result := create.AddKeyValueDoc(arguments[2], arguments[3], arguments[4])
		fmt.Println(result, isSuccess)

	} else if identifier == Others.Read {

	} else if identifier == Others.Update {

	} else if identifier == Others.Delete {

	} else {
		os.Exit(0)
	}

}
