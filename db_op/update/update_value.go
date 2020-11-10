package update

import (
	"effortlessdb/db_op"
	"effortlessdb/db_op/read"
	"encoding/json"
	"io/ioutil"
	"strings"
)

func UpdateValueByKey(collectionName string, _id string, key string, value string) (bool, []string) {
	isPresent, result := read.ReadKeyValueByIdKey(collectionName, _id, key)

	if !isPresent {
		return false, []string{}
	}

	in := []byte(result[0])
	var raw map[string]interface{}
	if err := json.Unmarshal(in, &raw); err != nil {
		//panic(err)
		return false, []string{}
	}
	raw[key] = value
	out, err := json.Marshal(raw)
	if err != nil {
		///panic(err)
		return false, []string{}

	}
	//println(string(out))

	path := collectionName + string(key[0]) + db_op.EfdExt
	read, err := ioutil.ReadFile(path)
	if err != nil {
		//panic(err)
		return false, []string{}
	}
	//fmt.Println(string(read))
	//fmt.Println(path)

	newContents := strings.Replace(string(read), result[0], string(out), -1)

	//fmt.Println(newContents)

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		return false, []string{}

		//		panic(err)
	}

	return true, strings.Fields(string(out))

}
