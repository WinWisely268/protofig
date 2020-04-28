protofig-proto:
	protoc -I protofig/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		-I . \
		--go_out=":protofig/" \
		--validate_out="lang=go:protofig/" \
		protofig/*.proto

api-proto:
	protoc \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		-I . \
		--go_out=":pkg" \
		--validate_out="lang=go:pkg" \
		api/*.proto
