package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	format "github.com/SAP/jenkins-library/pkg/format"
	fortify "github.com/SAP/jenkins-library/pkg/fortify"
	models "github.com/piper-validation/fortify-client-go/models"
)

func main() {
	fmt.Println("Piper FPR to SARIF converter standalone")
	fmt.Println("Run without SSC: ./converter result.fpr")
	fmt.Println("Run with SSC: ./converter result.fpr ServerURL AuthToken ProjectVersionID")
	if len(os.Args[1:]) == 0 {
		fmt.Println("need filename as first argument")
		return
	}
	fprFileName := os.Args[1]
	var sys fortify.System
	var sarif format.SARIF
	var err error
	if len(os.Args[2:]) == 3 {
		fmt.Println("SSC information found, attempting to run with audit data")
		sys = fortify.NewSystemInstance(os.Args[2], "/api/v1", os.Args[3], time.Minute*15)
		projectVersionID, _ := strconv.Atoi(os.Args[4])
		projectVersion := models.ProjectVersion{ID: int64(projectVersionID)}
		filterSet, _ := sys.GetFilterSetOfProjectVersionByTitle(projectVersion.ID, "SAP")
		fmt.Println("Running converter...")
		sarif, err = fortify.ConvertFprToSarif(sys, &projectVersion, fprFileName, filterSet)
	} else {
		fmt.Println("Not enough SSC information passed, running without audit data")
		fmt.Println("Syntax: ./converter result.fpr ServerURL AuthToken ProjectVersionID")
		fmt.Println("Running converter...")
		sarif, err = fortify.ConvertFprToSarif(sys, nil, fprFileName, nil)
	}
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
