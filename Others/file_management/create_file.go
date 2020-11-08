package file_management

import (
	"effortlessdb/Others"
	"fmt"
	"os"
)

func CreateFile(collectionName string) bool {

	file, err := os.Create(collectionName + Others.EfdExt)
	if err != nil {
		fmt.Println("Unable to create collection", collectionName)
		return false
	}

	defer file.Close()
	return true
}
