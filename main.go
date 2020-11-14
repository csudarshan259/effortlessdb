package main

import (
	"db_op"
	create2 "db_op/create"
	delete2 "db_op/delete"
	"db_op/read"
	"db_op/update"
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
				// isSuccess, result := read.ReadKeyValue(arguments[2], arguments[3])

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

		if len(os.Args) == 4 {

			isSuccess, result := delete2.DeleteByKey(arguments[2], arguments[3]) //collection_name key

			//fmt.Println("result",result)
			fmt.Println(result, isSuccess)

		} else if len(os.Args) == 5 {

			isSuccess, result := delete2.DeleteByKeyId(arguments[2], arguments[3], arguments[4]) //collection_name,_id,key
			fmt.Println(result, isSuccess)

		}

	} else {
		os.Exit(0)
	}

}
