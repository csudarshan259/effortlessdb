package create

import (
	others2 "effortlessdb/effortlessdb/others"
	fileManagement2 "effortlessdb/effortlessdb/others/fileManagement"
)

func AddKeyValueDoc(collectionName string, key string, value string) (bool, string) {
	if len(collectionName) > 24 {
		//fmt.Println(collectionName)
		return false, others2.InvalidCollectionName
	}
	isSuccess, result := fileManagement2.AddToFile(collectionName, key, value)
	//fmt.Println(result)
	return isSuccess, result
}
