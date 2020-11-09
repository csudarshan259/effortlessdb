package Object_Id_Generation

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateObjectId() string {

	var abData []byte = nil
	var strHex string
	abData = make([]byte, 16)
	rand.Read(abData)
	abData[6] = byte(0x40 | (int(abData[6]) & 0xf))
	abData[8] = byte(0x80 | (int(abData[8]) & 0x3f))

	strHex = hex.EncodeToString(abData)

	return strHex

}
