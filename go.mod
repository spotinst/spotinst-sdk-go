module github.com/spotinst/spotinst-sdk-go

go 1.20

require (
	github.com/stretchr/testify v1.8.4
	gopkg.in/ini.v1 v1.67.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v1.334.0 //incorrect versioning
)
