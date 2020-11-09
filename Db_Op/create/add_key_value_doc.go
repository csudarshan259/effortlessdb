package create

import (
	"effortlessdb/Db_Op"
	fileManagement2 "effortlessdb/Db_Op/Fiile_management"
)

func AddKeyValueDoc(collectionName string, key string, value string) (bool, string) {
	if len(collectionName) > 24 {
		//fmt.Println(collectionName)
		return false, Db_Op.InvalidCollectionName
	}
	isSuccess, result := fileManagement2.AddToFile(collectionName, key, value)
	//fmt.Println(result)
	return isSuccess, result
}
