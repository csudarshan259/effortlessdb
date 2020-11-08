package create

import (
	others2 "effortlessdb/effortlessdb/others"
	fileManagement2 "effortlessdb/effortlessdb/others/fileManagement"
)

func createCollection(collectionName string) (bool, string) {
	if len(collectionName) > 24 {

		return false, "Collection name cannot exceed 24 characters"
	}
	isCollectionCreated := fileManagement2.CreateFile(collectionName)
	if isCollectionCreated {
		return true, others2.Success
	}
	return false, others2.UnrecognisedError
}
