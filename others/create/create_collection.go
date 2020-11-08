package create

import (
	"effortlessdb"
	"effortlessdb/others/file_management"
)

func createCollection(collectionName string) (bool, string) {
	if len(collectionName) > 24 {

		return false, "Collection name cannot exceed 24 characters"
	}
	isCollectionCreated := file_management.CreateFile(collectionName)
	if isCollectionCreated {
		return true, effortlessdb.Success
	}
	return false, effortlessdb.UnrecognisedError
}
