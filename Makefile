.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: run
run:
	LOGXI=* go run .
