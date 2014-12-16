package main

import "fmt"

// import "os"
// import "os/exec"
// import "github.com/tj/go-search"
// import "github.com/tj/docopt"
// import "errors"
// import "encoding/json"
// import "path"
// import "net/http"

const Usage = `
This is a go package manage tool.

Usage:
	gpm install <package> [--save]
	gpm remove <package> [--save]

Arguments:
	init    --generator a package.json file.
	install --install packages define in the package.json file.
	search  --search package and list packages find in godoc.org.
	update  --update installed packages.

Options:
	-h|--help 		show this message.
	--version  		show version of current.	
`

const Version = `Version: gpm 0.0.1`

func main() {
	// args, _ := docopt.Parse(Usage, nil, true, Version, false)
	// println(args)
	res := search("github.com/tj/docpt")
	if len(res) == 0 {
		fmt.Sprintf("\033[31m")
		println("No packages find in godoc.org.")
	}
}
