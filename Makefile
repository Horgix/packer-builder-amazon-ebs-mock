GOAPP		= github.com/Horgix/packer-builder-amazon-ebs-mock
CWD		= `pwd`

SOURCES		= main.go \
		  amazon-ebs-mock/ \
		  packer-lib-mock/

build:: fmt
	go build -o packer-builder-amazon-ebs-mock

docker::
	docker run --rm \
	  -v "${CWD}":"/usr/src/${GOAPP}" \
	  -w "/usr/src/${GOAPP}" \
	  -e GOPATH=/usr/ \
	    golang:1.8 \
	      make

debug::
	PACKER_LOG=true packer build tests/simplest_packer.json

fmt::
	go fmt ${GO_PKGS}

test::
	go test ${GOAPP}/amazon-ebs-mock

readme_lint::
	docker run --rm \
	  -v ${CWD}/README.md:/data/README.md \
	  -v ${CWD}/lint_style.rb:/lint_style.rb \
	  mivok/markdownlint:0.4.0 --style /lint_style.rb .
