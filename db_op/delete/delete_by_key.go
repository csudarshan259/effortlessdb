package delete

import (
	"effortlessdb/db_op/update"
)

func DeleteByKey(collectionName string, key string) (bool, []string) {

	isSuccess, result := update.UpdateValueByKeyWithoutId(collectionName, key, "")

	if isSuccess == true {
		return isSuccess, result
	}

	return false, []string{}

}
