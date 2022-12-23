package main

import (
	"fmt"
	//"github.com/envoamr/Overcrawl/thread"
	"os"

	"github.com/envoamr/Overcrawl/web"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(dir)
	for i := 0; i <= 100; i++ {
		fmt.Println(web.SendMeta(dir + "/sample/txt"));
	}
}
