//go:generate statik -src=./tpl -f
package main

import (
	"fmt"
	"os"

	"github.com/rakyll/statik/fs"

	_ "github.com/FengGeSe/nest/statik"
)

func main() {

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	if err := fs.Walk(statikFS, "/", Show); err != nil {
		panic(err)
	}

	fmt.Println("end")
}

func Show(path string, info os.FileInfo, err error) error {

	fmt.Println(path, info.Name(), err)
	return nil
}
