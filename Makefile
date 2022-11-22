GO111MODULE:=on
GOBIN:=(go env GOPATH)/bin

go_install:
	go mod download && go generate && go build -o aoa ./ && mv ./aoa ~/.local/bin/
