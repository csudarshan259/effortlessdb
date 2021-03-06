package file_management

import (
	"db_op"
	"encoding/json"
	"log"
	"object_id_generation"
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
		db_op.Id: object_id_generation.GenerateObjectId(),
		key:      value,
	}

	//KVBytes := *(*[unsafe.Sizeof(kevValuePair)]byte)(unsafe.Pointer(&kevValuePair))

	b, err := json.Marshal(KeyValuePair)
	b = append(b, db_op.NewLine...)
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

	return true, db_op.Success
}
