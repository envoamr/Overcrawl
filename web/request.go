package web

import (
	"os"
	"github.com/envoamr/Overcrawl/misc"
)

var website string

func SendMeta(websiteName string) string {
	web, err := os.ReadFile(websiteName)
	website = string(web)

	if err != nil {
		panic(err)
	}

	startEnd := misc.Rand(2)
	var start int = int(startEnd[0])
	var off int = int(startEnd[1])
	return website[start:start+off]
}
