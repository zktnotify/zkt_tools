#!/bin/sh


protoc -I ./proto ./proto/*.proto --go_out=plugins=grpc:./pkg/protobuf

ls ./pkg/protobuf/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

echo "Done!"