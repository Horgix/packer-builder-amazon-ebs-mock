GOAPP		= github.com/Horgix/packer-builder-amazon-ebs-mock
CWD		= `pwd`

build:: fmt
	go build -o packer-builder-amazon-ebs-mock

docker::
	docker run --rm \
	  -v "${CWD}":"/usr/src/${GOAPP}" \
	  -w "/usr/src/${GOAPP}" \
	  -e GOPATH=/usr/ \
	    golang:1.8 \
	      make

fmt::
	gofmt -w main.go amazon-ebs-mock/

test::
	export PACKER_LOG=true
	packer build packer.json

readme_lint::
	docker run --rm \
	  -v ${CWD}/README.md:/data/README.md \
	  -v ${CWD}/lint_style.rb:/lint_style.rb \
	  mivok/markdownlint:0.4.0 --style /lint_style.rb .
