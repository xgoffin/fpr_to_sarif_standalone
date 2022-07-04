# FPR to SARIF standalone

## Usage

Without SSC:

Can be used locally, lightweight, no audit information

`go run main.go result.fpr`

With SSC:

Pull audit information from Fortify SSC

`go run .\main.go result.fpr ServerURL FortifyToken(PiperToken) ProjectVersionID`

Example:

`go run .\main.go result.fpr https://fortify-stage.tools.sap/ssc 11111111111111111111111111111111111111111111111 75485`

File can be found under `fortify/result.sarif`

## Compiling & running 

Seems to give faster results

`go build -o converter`

`./converter result.fpr`

`./converter result.fpr https://fortify-stage.tools.sap/ssc 11111111111111111111111111111111111111111111111 75485`

## Caveat

Running the converter twice in a row will replace the output file.