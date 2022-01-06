go-generate:
	@echo ">  Generating proto files..."
	@$(GOROOT)/bin/go generate -v ./main.go