package main

import "os/exec"
import "fmt"

func installAll() error {
	var cmd *exec.Cmd
	response, err := parse()
	if err != nil {
		return err
	}
	dependencies := response.(Request).Dependencies
	for _, v := range dependencies {
		cmd = exec.Command("go", "get", v)
		fmt.Println("installing package ", v)
		if err := cmd.Start(); err != nil {
			return err
		}
	}
	return nil
}
