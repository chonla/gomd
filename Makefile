test:
	go test ./... -run ^Test[^Integration]

integration-test:
	go test ./... -run ^TestIntegration

run:
	go run main.go