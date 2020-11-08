package fileManagement

import (
	others2 "effortlessdb/effortlessdb/others"
	"fmt"
	"os"
)

func CreateFile(collectionName string) bool {

	file, err := os.Create(collectionName + others2.EfdExt)
	if err != nil {
		fmt.Println("Unable to create collection", collectionName)
		return false
	}

	defer file.Close()
	return true
}
