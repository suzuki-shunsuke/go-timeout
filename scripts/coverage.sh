go test ./... -race -coverprofile=.coverage.txt -covermode=atomic || exit 1
go tool cover -html=.coverage.txt
