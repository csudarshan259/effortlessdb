package main

import (
	"effortlessdb/Db_Op"
	create2 "effortlessdb/Db_Op/create"
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
	if identifier == Db_Op.Create {

		isSuccess, result := create2.AddKeyValueDoc(arguments[2], arguments[3], arguments[4])
		fmt.Println(result, isSuccess)

	} else if identifier == Db_Op.Read {

	} else if identifier == Db_Op.Update {

	} else if identifier == Db_Op.Delete {

	} else {
		os.Exit(0)
	}

}
