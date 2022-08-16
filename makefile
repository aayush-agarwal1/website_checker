build:
	go clean
	go build -o build/website_checker -a website_checker/main.go

run:
	build/website_checker properties.yml

test:
	go test -v ./...

codeCoverage:
	go test -v -coverpkg=./... ./...

