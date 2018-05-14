mkdir pb
cd proto
protoc --go_out=plugins=grpc:../pb *