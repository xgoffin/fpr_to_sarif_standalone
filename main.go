package main

import (
	"fmt"
	"os"

	fortify "github.com/SAP/jenkins-library/pkg/fortify"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("need filename as first argument")
		return
	}
	fprFileName := os.Args[1]
	var sys fortify.System
	sarif, err := fortify.ConvertFprToSarif(sys, nil, nil, fprFileName, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fortify.WriteSarif(sarif)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Report written under ./fortify/result.sarif")
}
