module github.com/aiven/aiven-go-client/tools/exp

go 1.19

require (
	github.com/ChimeraCoder/gojson v1.1.0
	github.com/aiven/aiven-go-client v0.0.0
	github.com/google/go-cmp v0.5.8
	github.com/mitchellh/copystructure v1.2.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/pflag v1.0.5
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e
	golang.org/x/net v0.0.0-20220812174116-3211cb980234
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	golang.org/x/sys v0.0.0-20220817070843-5a390386f1f2 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/aiven/aiven-go-client => ../../
