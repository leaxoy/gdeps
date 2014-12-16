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

const (
	Usage = `
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

	Version = `Version: gpm 0.0.1`
)

func main() {
	// args, _ := docopt.Parse(Usage, nil, true, Version, false)
	// println(args)
	res := search("github.com/tj/docpt")
	if len(res) == 0 {
		fmt.Sprintf("\033[31m")
		println("No packages find in godoc.org.")
	}
}

// func read(s string) (*os.File, error) {
// 	_, err := os.Stat(s)
// 	if err != nil {
// 		return nil, errors.New("package.json not exist.please check or modify it.")
// 	}
// 	file, err := os.Open(s)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return file, nil
// }

// func parse() map[string]interface{} {
// 	s := path.Join(dir(), "package.json")
// 	f, err := read(s)
// 	if err != nil {
// 		fmt.Println(fmt.Sprintf("\033[32m%s", "can't find package.json file, please modify it."))
// 		return nil
// 	}
// 	var args Arguments
// 	result := make(map[string]interface{})
// 	json.NewDecoder(f).Decode(&args)
// 	for _, arg := range args.Arg {
// 		result[arg.key] = arg.val
// 	}
// 	return result
// }

// func dir() string {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		return ""
// 	}
// 	return dir
// }
