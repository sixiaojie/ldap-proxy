package authorization

import (
	"encoding/base64"
)

func EntryBase64(info string)(string){
	info = info +":" + info
	base64string :=  base64.StdEncoding.EncodeToString([]byte(info))
	return "Basic "+base64string
}

