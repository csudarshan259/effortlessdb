package create

import (
	"effortlessdb/Db_Op"
	fileManagement2 "effortlessdb/Db_Op/Fiile_management"
)

func createCollection(collectionName string) (bool, string) {
	if len(collectionName) > 24 {

		return false, "Collection name cannot exceed 24 characters"
	}
	isCollectionCreated := fileManagement2.CreateFile(collectionName)
	if isCollectionCreated {
		return true, Db_Op.Success
	}
	return false, Db_Op.UnrecognisedError
}
