build:: fmt
	go build -o packer-builder-amazon-ebs-mock

fmt::
	gofmt -w main.go amazon-ebs-mock/

test::
	export PACKER_LOG=true
	packer build packer.json
