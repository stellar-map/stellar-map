CMDS := stellarmap migrate

$(CMDS):
	go build -o bin/$@ go/cmd/$@/*.go

dev:
	go build -o bin/stellarmap go/cmd/stellarmap/*.go
	bin/stellarmap

test:
	ginkgo -p go/...

lint:
	gometalinter go/...

tools:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/onsi/ginkgo/ginkgo
	go get -u github.com/alecthomas/gometalinter

.PHONY: $(CMDS) dev test lint tools
