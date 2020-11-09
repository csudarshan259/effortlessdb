package File_management

import (
	"effortlessdb/Db_Op"
	"effortlessdb/Object_Id_Generation"
	"encoding/json"
	"log"
	"os"
)

type KVPair map[string]string

func AddToFile(collectionName string, key string, value string) (bool, string) {
	_, fileName := CreateFileIFNotExists(collectionName, string(key[0]))
	//fmt.Println("filename1",fileName)

	file, err := os.OpenFile(fileName,
		//	os.O_APPEND|os.O_APPEND|os.O_WRONLY, 0644)
		os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	KeyValuePair := KVPair{
		Db_Op.Id: Object_Id_Generation.GenerateObjectId(),
		key:      value,
	}

	//KVBytes := *(*[unsafe.Sizeof(kevValuePair)]byte)(unsafe.Pointer(&kevValuePair))

	b, err := json.Marshal(KeyValuePair)
	if err != nil {
		//		fmt.Println("Unable to marshal")
		return false, "Unable to marshal key value pair"
	}
	//	fmt.Println(KeyValuePair)
	_, err = file.Write(b)

	if err != nil {
		//		fmt.Println("Unable to write")
		return false, "Unable to write to file"
	}

	return true, Db_Op.Success
}
