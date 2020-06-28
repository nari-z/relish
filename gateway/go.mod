module github.com/nari-z/relish/gateway

go 1.14

replace github.com/nari-z/relish => ../
replace github.com/nari-z/relish/gateway => ./
replace github.com/nari-z/relish/authneticator => ../authenticator

require (
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/loads v0.19.5
	github.com/go-openapi/runtime v0.19.19
	github.com/go-openapi/spec v0.19.8
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/jessevdk/go-flags v1.4.0
	github.com/nari-z/relish v0.0.0-20200607135311-fb82e1fe4d89
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	google.golang.org/grpc v1.30.0
)
