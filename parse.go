package main

import "path"
import "os"
import "encoding/json"

type Request struct {
	Name         string
	Version      string
	Author       string
	Email        string
	Homepage     string
	Dependencies Dependencies
}

type Dependencies []string

func parse() (interface{}, error) {
	pwd, _ := os.Getwd()
	var req Request
	pj := path.Join(pwd, "package.json")
	file := read(pj)
	err := json.NewDecoder(file).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func read(s string) *os.File {
	_, err := os.Stat(s)
	if err != nil {
		return nil
	}
	file, err := os.Open(s)
	if err != nil {
		return nil
	}
	return file
}

func Init() {

}
