build:
	go build -o build/website_checker website_checker/main.go

run:
	build/website_checker properties.yml

test:
	go test -v ./...

codeCoverage:
	go test -v -coverpkg=./... ./...

