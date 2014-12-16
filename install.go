package main

import "os/exec"

func install() error {
	var cmd *exec.Cmd
	dependencies := parse().Dependencies
	for _, v := range dependencies {
		cmd = exec.Command("go", "get", v)
		if err := cmd.Start(); err != nil {
			return err
		}
	}
	return nil
}
