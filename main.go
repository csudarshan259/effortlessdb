package main

import (
	others2 "effortlessdb/effortlessdb/others"
	create2 "effortlessdb/effortlessdb/others/create"
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
	if identifier == others2.Create {

		isSuccess, result := create2.AddKeyValueDoc(arguments[2], arguments[3], arguments[4])
		fmt.Println(result, isSuccess)

	} else if identifier == others2.Read {

	} else if identifier == others2.Update {

	} else if identifier == others2.Delete {

	} else {
		os.Exit(0)
	}

}
