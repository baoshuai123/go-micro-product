

.PHONY: proto
proto:
	docker run --rm  -v d:/GOLANG/src/taobao/product:/d/GOLANG/src/taobao/product -w /d/GOLANG/src/taobao/product  -e ICODE=2606C833CD172F4C cap1573/cap-protoc -I ./ --micro_out=./ --go_out=./ ./proto/product/product.proto

.PHONY: build
build:
	set CGO_ENABLED=0
	set GOOS=linux
	set GOARCH=amd64
	go build -o product-service

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t product-service:latest
