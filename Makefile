TEST?=$$(GO15VENDOREXPERIMENT=1 go list ./... | grep -v /vendor/)
VETARGS?=-asmdecl -atomic -bool -buildtags -copylocks -methods -nilfunc -printf -rangeloops -shift -structtags -unsafeptr

default: test

test:
	TF_ACC= GO15VENDOREXPERIMENT=1 go test -v $(TEST) $(TESTARGS) -timeout=30s -parallel=4

fmt:
	gofmt -w .
