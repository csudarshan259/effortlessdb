package update

import (
	"db_op"
	"db_op/read"
	"encoding/json"
	"fmt"
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

//function only for delete purpose
func UpdateValueByKeyWithoutId(collectionName string, key string, value string) (bool, []string) {
	isPresent, result := read.ReadKeyValue(collectionName, key)
	var emptyStringSlice []string

	//fmt.Println("result",result)
	if !isPresent {
		//fmt.Println("notpresent",result)
		return false, emptyStringSlice
	}
	fmt.Println(result)
	var newContents string
	for i := 0; i < len(result); i++ {

		result[i] = strings.TrimSpace(result[i])

		if result[i] == "" {
			continue
		}

		in := []byte(result[i])
		var raw map[string]interface{}

		if err := json.Unmarshal(in, &raw); err != nil {
			//fmt.Println(i,err)
			return false, emptyStringSlice
		}
		println(raw)
		raw[key] = value
		_, err := json.Marshal("")

		if err != nil {
			//panic(err)
			return false, emptyStringSlice

		}
		//println("out",string(out))

		path := collectionName + string(key[0]) + db_op.EfdExt
		read, err := ioutil.ReadFile(path)
		if err != nil {
			//	panic(err)
			//fmt.Println("nul")

			return false, emptyStringSlice
		}
		//fmt.Println("read",string(read))
		//fmt.Println(path)

		//newContents := strings.Replace(string(read), result[0], string(out), -1)

		//fmt.Println(result[i])
		newContents = strings.Replace(string(read), result[i]+"\n", "", -1)
		//newContents = strings.Replace(newContents, "\n","", -1)
		//newContents = strings.Replace(newContents, "}","}\n", -1)

		//newContents = newContents+"\n"

		//fmt.Println(newContents)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			//fmt.Println("nul")
			//		panic(err)
			return false, emptyStringSlice
		}

	}
	//fmt.Println("finalresult",result)
	return true, result
}
