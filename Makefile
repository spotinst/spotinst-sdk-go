TEST?=$$(GO15VENDOREXPERIMENT=1 go list ./... | grep -v /vendor/)
VETARGS?=-asmdecl -atomic -bool -buildtags -copylocks -methods -nilfunc -printf -rangeloops -shift -structtags -unsafeptr

default: test

# test runs the unit tests
test:
	TF_ACC= GO15VENDOREXPERIMENT=1 go test $(TEST) $(TESTARGS) -timeout=30s -parallel=4

# install installs all the dependencies using Glide (https://github.com/Masterminds/glide).
install:
	GO15VENDOREXPERIMENT=1 glide install
