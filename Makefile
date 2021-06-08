.PHONY: getdeps protos gorelease

getdeps:
	which protoc
	export GO111MODULE=on
	go get -u -v google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u -v google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u -v github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install -v google.golang.org/protobuf/cmd/protoc-gen-go
	go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install -v github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

protos:
	protoc -I ./proto/ -I third_party/googleapis -I third_party/grpc-gateway \
    --openapiv2_out=./openapiv2 --openapiv2_opt logtostderr=true \
    --go_out=. --go_opt=module=gitlab.com/insanitywholesale/lister \
    --go-grpc_out=. --go-grpc_opt=module=gitlab.com/insanitywholesale/lister \
    --grpc-gateway_out ./proto \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
	proto/v1/*.proto

gorelease:
	go install -v github.com/goreleaser/goreleaser@latest
	goreleaser --snapshot --skip-publish --rm-dist
