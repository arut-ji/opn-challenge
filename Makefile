FILE_PATH=./data/fng.1000.csv

run:
	go run -race main.go -publicKey=$(OMISE_PUBLIC_KEY) -secretKey=$(OMISE_SECRET_KEY) $(FILE_PATH)

build:
	go build -o bin/go-tamboon

clean:
	go clean && rm -rf bin/