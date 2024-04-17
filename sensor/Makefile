
bin/customsensor: *.go cmd/module/*.go go.*
	go build -o bin/customsensor cmd/module/cmd.go

bin/remoteserver: *.go cmd/remote/*.go go.*
	go build -o bin/remoteserver cmd/remote/cmd.go

lint:
	gofmt -w -s .

updaterdk:
	go get go.viam.com/rdk@latest
	go mod tidy

module: bin/customsensor
	tar czf module.tar.gz bin/customsensor