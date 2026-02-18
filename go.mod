module github.com/gissleh/litxap

go 1.24

toolchain go1.24.13

require github.com/stretchr/testify v1.9.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//for testing on a local machine's fwew-lib
//replace github.com/fwew/fwew-lib/v5 => ../fwew-lib
