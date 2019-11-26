//go:generate statik -src=./tpl -f
package main

import (
	"fmt"
	"os"

	"github.com/rakyll/statik/fs"

	cmd "github.com/FengGeSe/nest/cmd"
	_ "github.com/FengGeSe/nest/statik"
)

var app = cmd.App

var (
	_version   = "unknown"
	_gitCommit = "unknown"
	_goVersion = "unknown"
	_buildTime = "unknown"
	_osArch    = "unknown"
)

func init() {
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr
	app.Version = _version
	app.HideVersion = true
	app.Description = fmt.Sprintf(`Version:          %s
   Go version:       %s
   Git commit:       %s
   Built:            %s
   OS/Arch:          %s`, _version, _goVersion, _gitCommit, _buildTime, _osArch)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func WalkStatikFs() {
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
