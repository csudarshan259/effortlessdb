package file_management

import (
	"effortlessdb/db_op"
	"fmt"
	"os"
)

func CreateFile(collectionName string) bool {

	file, err := os.Create(collectionName + db_op.EfdExt)
	if err != nil {
		fmt.Println("Unable to create collection", collectionName)
		return false
	}

	defer file.Close()
	return true
}
