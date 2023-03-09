package main

import (
//	"flag"
	"fmt"
//	"os"
//	"os/exec"
)

func main() {
	// Place your code here.
	env, err := ReadDir("testdata/env/")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, val := range env {
		fmt.Printf("Value=%s  flag=%v\n", val.Value, val.NeedRemove)
	}
}
