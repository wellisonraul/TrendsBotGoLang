build-local:
	go build -o local

serve-local: build-local
	./local

build-linux:
	GOOS=linux GOARCH=386 go build -a --ldflags="-s"

deploy: build-linux
	gcloud compute scp ./in0984 instance-1:~/
