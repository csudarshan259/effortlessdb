package main

import (
	"effortlessdb/db_op"
	create2 "effortlessdb/db_op/create"
	"effortlessdb/db_op/read"
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
	if identifier == db_op.Create {

		isSuccess, result := create2.AddKeyValueDoc(arguments[2], arguments[3], arguments[4])
		fmt.Println(result, isSuccess)

	} else if identifier == db_op.Read {

		isSuccess, result := read.ReadKeyValue(arguments[2], arguments[3])
		fmt.Println(result, isSuccess)

	} else if identifier == db_op.Update {

	} else if identifier == db_op.Delete {

	} else {
		os.Exit(0)
	}

}
