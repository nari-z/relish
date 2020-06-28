module github.com/nari-z/relish

go 1.14

replace github.com/nair-z/relish/relish => ./

replace github.com/nair-z/relish/gateway => ./gateway

replace github.com/nair-z/relish/authneticator => ./authenticator

require (
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/loads v0.19.5
	github.com/go-openapi/runtime v0.19.19
	github.com/go-openapi/spec v0.19.8
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/golang/protobuf v1.4.2
	github.com/jessevdk/go-flags v1.4.0
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)
