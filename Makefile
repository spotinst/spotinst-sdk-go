.PHONY: test
test: fmtcheck
	go test -i $$(go list ./... | grep -v 'vendor') $(TESTARGS) -timeout=30s -parallel=4

.PHONY: depinit
depinit:
	dep init

.PHONY: depensure
depensure:
	dep ensure

.PHONY: fmt
fmt:
	gofmt -w $$(find . -name '*.go' | grep -v vendor)

.PHONY: fmtcheck
fmtcheck:
	@! gofmt -d $$(find . -name '*.go' | grep -v vendor) | grep '^'

.PHONY: vet
vet:
	go tool vet -all -v $$(find . -name '*.go' | grep -v vendor)