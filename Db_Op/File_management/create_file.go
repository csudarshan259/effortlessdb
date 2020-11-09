package File_management

import (
	"effortlessdb/Db_Op"
	"fmt"
	"os"
)

func CreateFile(collectionName string) bool {

	file, err := os.Create(collectionName + Db_Op.EfdExt)
	if err != nil {
		fmt.Println("Unable to create collection", collectionName)
		return false
	}

	defer file.Close()
	return true
}
