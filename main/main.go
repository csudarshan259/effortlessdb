package main

import (
	"effortlessdb"
	create2 "effortlessdb/others/create"
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
	if identifier == effortlessdb.Create {

		isSuccess, result := create2.AddKeyValueDoc(arguments[2], arguments[3], arguments[4])
		fmt.Println(result, isSuccess)

	} else if identifier == effortlessdb.Read {

	} else if identifier == effortlessdb.Update {

	} else if identifier == effortlessdb.Delete {

	} else {
		os.Exit(0)
	}

}
