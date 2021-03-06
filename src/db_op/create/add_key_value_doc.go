package create

import (
	"db_op"
	fileManagement2 "db_op/file_management"
)

func AddKeyValueDoc(collectionName string, key string, value string) (bool, string) {
	if len(collectionName) > 24 {
		//fmt.Println(collectionName)
		return false, db_op.InvalidCollectionName
	}
	isSuccess, result := fileManagement2.AddToFile(collectionName, key, value)
	//fmt.Println(result)
	return isSuccess, result
}
