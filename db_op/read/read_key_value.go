package read

import (
	"effortlessdb/db_op"
	"io/ioutil"
	"strings"
)

func ReadKeyValue(collectionName string, key string) (bool, []string) {

	data, err := ioutil.ReadFile(collectionName + (string([]rune(key)[0])) + db_op.EfdExt)

	if err != nil {
		//fmt.Println(err)
		return false, nil
	}

	splitData := strings.Split(string(data), "\n")
	splitDataLength := len(splitData)
	result := make([]string, splitDataLength)

	j := 0
	for i := 0; i < len(splitData)-1; i++ {
		//fmt.Println(splitData[i])
		var splitData1 = strings.Split(splitData[i], ",")
		var splitData2 = strings.Split(splitData1[1], ":")
		//fmt.Println(splitData2[0])
		var trimSplit = strings.Replace(splitData2[0], "\"", "", -1)
		//fmt.Println(key)
		if string(trimSplit) == key {
			result[j] = splitData[i]
			//	fmt.Println(result[j])
			j++
		}

	}
	if j == 0 {
		return false, nil
	}
	return true, result

}
