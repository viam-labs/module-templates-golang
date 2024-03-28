bin/custompowersensor: *.go cmd/module/*.go go.*
	go build -o bin/custompowersensor cmd/module/cmd.go

bin/remoteserver: *.go cmd/remote/*.go go.*
	go build -o bin/remoteserver cmd/remote/cmd.go

lint:
	gofmt -w -s .

updaterdk:
	go get go.viam.com/rdk@latest
	go mod tidy

module: bin/custompowersensor
	tar czf module.tar.gz bin/custompowersensor