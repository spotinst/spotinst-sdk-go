TEST?=$$(GO15VENDOREXPERIMENT=1 go list ./... | grep -v /vendor/)
VETARGS?=-all

default: test

# test runs the unit tests
test:
	TF_ACC= GO15VENDOREXPERIMENT=1 go test -v $(TEST) $(TESTARGS) -timeout=30s -parallel=4

# fmt runs the Go format tool `gofmt` to format the code
fmt:
	gofmt -w .

# vet runs the Go source code static analysis tool `vet` to find any common errors.
vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo "go tool vet $(VETARGS) ."
	@GO15VENDOREXPERIMENT=1 go tool vet $(VETARGS) $$(ls -d */ | grep -v vendor) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi