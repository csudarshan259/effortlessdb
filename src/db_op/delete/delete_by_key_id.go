package delete

import (
	"db_op"
	"db_op/read"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func DeleteByKeyId(collectionName string, _id string, key string) (bool, []string) {

	isPresent, result := read.ReadKeyValueByIdKey(collectionName, _id, key)
	value := ""

	//fmt.Println("result",result)
	if !isPresent {
		return false, []string{}
	}

	var newContents string
	for i := 0; i < len(result); i++ {

		result[i] = strings.TrimSpace(result[i])

		if result[i] == "" {
			continue
		}

		in := []byte(result[i])
		var raw map[string]interface{}
		if err := json.Unmarshal(in, &raw); err != nil {
			//panic(err)
			fmt.Println("unmarshal")
			return false, []string{}
		}
		//raw[key] = value
		raw[_id] = value
		//_, err := json.Marshal("")
		//if err != nil {
		//	//panic(err)
		//	//return false, []string{}
		//
		//}
		//println("out",string(out))

		path := collectionName + string(key[0]) + db_op.EfdExt
		read, err := ioutil.ReadFile(path)
		if err != nil {
			//panic(err)
			//return false, []string{}
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
			//fmt.Println("write file")
			return false, []string{}

			//panic(err)
		}

	}
	//	fmt.Println("finalresult",result)

	return true, result

}
