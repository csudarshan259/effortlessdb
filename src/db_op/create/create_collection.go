package create

import (
	"db_op"
	fileManagement2 "db_op/file_management"
)

func createCollection(collectionName string) (bool, string) {
	if len(collectionName) > 24 {

		return false, "Collection name cannot exceed 24 characters"
	}
	isCollectionCreated := fileManagement2.CreateFile(collectionName)
	if isCollectionCreated {
		return true, db_op.Success
	}
	return false, db_op.UnrecognisedError
}
