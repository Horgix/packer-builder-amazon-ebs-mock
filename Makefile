build:: fmt
	go build -o packer-builder-amazon-ebs-mock

fmt::
	go fmt

test::
	export PACKER_LOG=true
	packer build packer.json
