package file_management

import (
	"effortlessdb/main/others"
	"os"
)

func CreateFileIFNotExists(collectionName string, _character string) (bool, string) {
	fileName := collectionName + _character + others.EfdExt
	//	fmt.Println("filename",fileName)
	_, err := os.Stat(fileName)
	isFileExist := !os.IsNotExist(err)
	//println(os.IsNotExist(err),isFileExist)
	if isFileExist {
		//		fmt.Println("File exists")
		return true, fileName
	}

	file, err := os.Create(fileName)
	if err != nil {
		//		fmt.Println("Unable to create collection", collectionName)
		return false, fileName
	}
	defer file.Close()
	return true, fileName
}
