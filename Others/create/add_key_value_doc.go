package create

import (
	"effortlessdb/Others"
	"effortlessdb/Others/file_management"
)

func AddKeyValueDoc(collectionName string, key string, value string) (bool, string) {
	if len(collectionName) > 24 {
		//fmt.Println(collectionName)
		return false, Others.InvalidCollectionName
	}
	isSuccess, result := file_management.AddToFile(collectionName, key, value)
	//fmt.Println(result)
	return isSuccess, result
}
