# frontend

Run the following command to restore dependencies to `vendor/` directory:

    dep ensure --vendor-only

```
apt-get install -y protoc-gen-go
```


how i build demo.pb.go
```
docker run -it golang:1.16-buster /bin/bash
apt-get update
apt-get install -y protobuf-compiler vim nano
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
mkdir genproto
nano demo.proto
protoc  --go_out=plugins=grpc:genproto -I./ demo.proto
ls genproto/demo.pb.go

```
