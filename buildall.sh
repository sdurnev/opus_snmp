GOROOT=/Users/svdu/go/go1.13.1
GOPATH=/Users/svdu/go
VERSION=0.00.1

GOOS=linux
GOARCH=arm
GOARM=5
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/opus_snmp/bin/$VERSION/arm/opus_snmp /Users/svdu/GolandProjects/opus_snmp/main.go
GOOS=windows
GOARCH=amd64
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/opus_snmp/bin/$VERSION/windows/opus_snmp /Users/svdu/GolandProjects/opus_snmp/main.go
GOOS=linux
GOARCH=amd64
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/opus_snmp/bin/$VERSION/linux/opus_snmp /Users/svdu/GolandProjects/opus_snmp/main.go
GOOS=darwin
GOARCH=amd64
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/opus_snmp/bin/$VERSION/macos/opus_snmp /Users/svdu/GolandProjects/opus_snmp/main.go