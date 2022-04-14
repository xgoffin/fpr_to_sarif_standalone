# FPR to SARIF standalone

## Usage

`go run main.go result.fpr`

File can be found under `fortify/result.sarif`

## Compiling & running 

Seems to give faster results

`go build -o converter`
`./converter result.fpr`

## Caveat

Running the converter twice in a row will replace the output file.