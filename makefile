build:
	go clean
	go build -o build/website_checker -a website_checker/main.go

run:
	build/website_checker properties.yml

test:
	go test -v ./...

codeCoverage:
	go test -v -coverpkg=./... ./...

.PHONY: build

postDummyData:
	curl --location --request POST '127.0.0.1:8080/websites' \
    --header 'Content-Type: application/json' \
    --data-raw '{ \
        "websites": [ \
            "www.google.com", \
            "www.facebook.com", \
            "www.fakewebsite1.com", \
            "www.youtube.com", \
            "www.yahoo.com", \
            "www.bing.com" \
        ] \
    }'

postWrongData:
	curl --location --request POST '127.0.0.1:8080/websites' \
        --header 'Content-Type: application/json' \
        --data-raw '{ \
            "websites": [ \
                "invalidURI" \
            ] \
        }'

getData1:
	curl --location --request GET '127.0.0.1:8080/websites'

getData2:
	curl --location --request GET '127.0.0.1:8080/websites?name=www.google.com'

getData3:
	curl --location --request GET '127.0.0.1:8080/websites?name=www.facebook.com,www.fakefakewebsite1.com'

getData4:
	curl --location --request GET '127.0.0.1:8080/websites?name=invalidURI,newURI,www.unseenURI.com'