GOPKG		= "github.com/Horgix/packer-builder-amazon-ebs-mock"
GO_PKGS		= `go list ./... | grep -v /vendor/`
CWD		= `pwd`

build:: fmt
	go build -o packer-builder-amazon-ebs-mock

docker::
	docker run --rm \
	  -v "${CWD}":"/usr/src/${GOPKG}" \
	  -w "/usr/src/${GOPKG}" \
	  -e GOPATH=/usr/ \
	    golang:1.8 \
	      make

debug::
	PACKER_LOG=true packer build tests/simplest_packer.json

fmt::
	go fmt ${GO_PKGS}

test::
	go test ${GO_PKGS}


readme_lint::
	docker run --rm \
	  -v ${CWD}/README.md:/data/README.md \
	  -v ${CWD}/lint_style.rb:/lint_style.rb \
	  mivok/markdownlint:0.4.0 --style /lint_style.rb .
