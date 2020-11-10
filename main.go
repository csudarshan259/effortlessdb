package main

import (
	"effortlessdb/db_op"
	create2 "effortlessdb/db_op/create"
	"effortlessdb/db_op/read"
	"effortlessdb/db_op/update"
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args

	if len(arguments) == 0 {
		os.Exit(0)
	}

	//println(len(os.Args))

	identifier := strings.ToLower(arguments[1])
	//fmt.Println(identifier)
	if identifier == db_op.Create {

		isSuccess, result := create2.AddKeyValueDoc(arguments[2], arguments[3], arguments[4])
		fmt.Println(result, isSuccess)

	} else if identifier == db_op.Read {

		if len(os.Args) == 4 {

			if strings.Split(arguments[3], ":")[0] == "key" {
				key := strings.Split(arguments[3], ":")[1]

				isSuccess, result := read.ReadKeyValue(arguments[2], key)
				//				isSuccess, result := read.ReadKeyValue(arguments[2], arguments[3])

				fmt.Println(result, isSuccess)
			}
			//else if strings.Split(arguments[3],":")[0]=="_id" {
			//	_id := strings.Split(arguments[3],":")[1]
			//	println(_id)
			//	isSuccess, result := read.ReadKeyValueById(arguments[2], _id)
			//	fmt.Println(result, isSuccess)
			//
			//}

		}
		if len(os.Args) == 5 {
			//if strings.Split(arguments[3],":")[0]=="_id"{
			isSuccess, result := read.ReadKeyValueByIdKey(arguments[2], strings.Split(arguments[3], ":")[1], strings.Split(arguments[4], ":")[1])
			//}

			fmt.Println(result, isSuccess)
		}

	} else if identifier == db_op.Update {
		isSuccess, result := update.UpdateValueByKey(arguments[2], strings.Split(arguments[3], ":")[1], strings.Split(arguments[4], ":")[1], strings.Split(arguments[5], ":")[1])

		fmt.Println(result, isSuccess)

	} else if identifier == db_op.Delete {

	} else {
		os.Exit(0)
	}

}
