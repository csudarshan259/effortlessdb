package fileManagement

import (
	others2 "effortlessdb/effortlessdb/others"
	"os"
)

func CreateFileIFNotExists(collectionName string, _character string) (bool, string) {
	fileName := collectionName + _character + others2.EfdExt
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
