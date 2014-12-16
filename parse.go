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

func parse() Request {
	pwd, _ := os.Getwd()
	var req Request
	pj := path.Join(pwd, "")
	file := read(pj)
	err := json.NewDecoder(file).Decode(&req)
	if err != nil {
		return req
	}
	return req
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
