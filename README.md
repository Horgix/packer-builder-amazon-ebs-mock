# Packer Builder - Amazon EBS Mock

[![Build Status](https://travis-ci.org/Horgix/packer-builder-amazon-ebs-mock.svg?branch=master)](https://travis-ci.org/Horgix/packer-builder-amazon-ebs-mock)
[![Coverage Status](https://coveralls.io/repos/github/Horgix/packer-builder-amazon-ebs-mock/badge.svg?branch=master)](https://coveralls.io/github/Horgix/packer-builder-amazon-ebs-mock?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/horgix/packer-builder-amazon-ebs-mock)](https://goreportcard.com/report/github.com/horgix/packer-builder-amazon-ebs-mock)

## What is this

**This Packer plugin is a Builder that mocks the [official Amazon EBS
builder](https://www.packer.io/docs/builders/amazon-ebs.html)**.

Why do you want that? Probably for testing purpose! At least this is my case,
and it was built to help me test my [Packer post-processor that sends Slack
notifications](https://github.com/Horgix/packer-post-processor-slack-notifications);
I **really** don't want to wait for a real AMI to build everytime each time I
want to run a test...

## How it does it

In order to generate something realist, it will currently generate some random
AMI names based on existing AWS regions but with `ami-12345678` as ID, just to
make sure you notice it fast if for whatever reason you end up thinking you're
running a true build.

| Option    | Description | Default |
| :-------: | :---------: | :-----: |
| Generate  | bla         | False   |
| Amount    | bla         | 1       |

## How to use it

Basically refer to the [Packer documentation on building
images](https://www.packer.io/intro/getting-started/build-image.html) about how
to call it in your `packer.json`, since it's working as any builder. To learn
how to **install** this plugin, refer to the [Packer documentation on
installing
plugins](https://www.packer.io/docs/extending/plugins.html#installing-plugins)

## How to build it

```shell
# Get dependencies
glide install
# Build the binary
make
```

## Example

### Most simple Packer JSON input (available in `tests/simplest_packer.json`)

```json
{
  "builders": [
    {
      "type":           "amazon-ebs-mock"
    }
  ]
}
```

### Output of such a build

`packer build ./tests/simplest_packer.json`

```raw
amazon-ebs-mock output will be in this color.

==> amazon-ebs-mock: I'm doing nothing...
==> amazon-ebs-mock: Nothing done with success!
==> amazon-ebs-mock: Generating mock Artifact...
==> amazon-ebs-mock: Generated mock Artifact with success :)
Build 'amazon-ebs-mock' finished.

==> Builds finished. The artifacts of successful builds are:
--> amazon-ebs-mock: AMIs were created:

ap-northeast-1: ami-12345678
ap-northeast-2: ami-12345678
ca-central-1: ami-12345678
eu-west-1: ami-12345678
eu-west-2: ami-12345678
sa-east-1: ami-12345678
us-west-2: ami-12345678
```

### Duration

If you time this build you'll end up with this:

```raw
packer build tests/simplest_packer.json  0.02s user 0.00s system 73% cpu 0.036 total
```

Which is all the point of this mock :)

### Want more details? Enable Packer logging!

See at the end of this README, it's huge :)

## Misc informations

- Since this is a small project that I'm also using as a test purpose, I went
  with `glide` as a go dependencies manager. If you feel more comfortable with
  `godep`, feel free to file a Pull Request replacing glide by godep and I'll
  happily accept it :)
- The `Makefile` currently provides 3 rules:
    - `build` (default rule): build a binary of the plugin called
      `packer-builder-amazon-ebs-mock`
    - `fmt`: format this plugin code according to Go standards
    - `test`: for now it just builds a packer.json...
- The code is not really commented, but feel free to open a issue or [contact
  me](TODO LINK) by any mean if you have some questions or want explanations
  about it

### Samples of `manifest.json`

Produced by the `manifest` post-processor with this builder

```json
{
  "builds": [
    {
      "name": "amazon-ebs-mock",
      "builder_type": "amazon-ebs-mock",
      "build_time": 1501605111,
      "files": null,
      "artifact_id": "ap-northeast-1:ami-12345678,ap-south-1:ami-12345678,ap-southeast-2:ami-12345678,eu-central-1:ami-12345678,sa-east-1:ami-12345678,us-east-1:ami-12345678,us-west-1:ami-12345678",
      "packer_run_uuid": "0c426d21-4e6a-6f56-e9aa-c98d2d325c65"
    },
    {
      "name": "amazon-ebs-mock",
      "builder_type": "amazon-ebs-mock",
      "build_time": 1501679363,
      "files": null,
      "artifact_id": "ap-northeast-1:ami-12345678,ap-northeast-2:ami-12345678,ap-southeast-1:ami-12345678,ca-central-1:ami-12345678,eu-west-2:ami-12345678,us-east-2:ami-12345678,us-west-1:ami-12345678,us-west-2:ami-12345678",
      "packer_run_uuid": "a6aeefb5-2083-3721-8b0c-b9a1c51608a7"
    }
  ],
  "last_run_uuid": "a6aeefb5-2083-3721-8b0c-b9a1c51608a7"
}
```

## Details of a basic build

`PACKER_LOG=true packer build -color=false ./tests/simplest_packer.json`

```raw
[INFO] Packer version: 0.12.3
Packer Target OS/Arch: linux amd64
Built with Go Version: go1.8
Detected home directory from env var: /home/horgix
[DEBUG] Discovered plugin: amazon-ebs-mock = MY_GOPATH/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock
Using internal plugin for vmware-vmx
Using internal plugin for amazon-instance
Using internal plugin for azure-arm
Using internal plugin for digitalocean
Using internal plugin for openstack
Using internal plugin for parallels-iso
Using internal plugin for virtualbox-iso
Using internal plugin for virtualbox-ovf
Using internal plugin for amazon-chroot
Using internal plugin for amazon-ebsvolume
Using internal plugin for googlecompute
Using internal plugin for oneandone
Using internal plugin for parallels-pvm
Using internal plugin for profitbricks
Using internal plugin for vmware-iso
Using internal plugin for amazon-ebs
Using internal plugin for docker
Using internal plugin for null
Using internal plugin for triton
Using internal plugin for amazon-ebssurrogate
Using internal plugin for cloudstack
Using internal plugin for file
Using internal plugin for hyperv-iso
Using internal plugin for qemu
Using internal plugin for ansible
Using internal plugin for ansible-local
Using internal plugin for windows-restart
Using internal plugin for windows-shell
Using internal plugin for converge
Using internal plugin for file
Using internal plugin for powershell
Using internal plugin for puppet-masterless
Using internal plugin for puppet-server
Using internal plugin for shell-local
Using internal plugin for chef-client
Using internal plugin for chef-solo
Using internal plugin for salt-masterless
Using internal plugin for shell
Using internal plugin for artifice
Using internal plugin for atlas
Using internal plugin for amazon-import
Using internal plugin for checksum
Using internal plugin for compress
Using internal plugin for googlecompute-export
Using internal plugin for docker-tag
Using internal plugin for manifest
Using internal plugin for shell-local
Using internal plugin for docker-import
Using internal plugin for docker-push
Using internal plugin for docker-save
Using internal plugin for vagrant
Using internal plugin for vagrant-cloud
Using internal plugin for vsphere
Detected home directory from env var: /home/horgix
Attempting to open config file: /home/horgix/.packerconfig
[WARN] Config file doesn't exist: /home/horgix/.packerconfig
Packer config: &{DisableCheckpoint:false DisableCheckpointSignature:false PluginMinPort:10000 PluginMaxPort:25000 Builders:map[amazon-ebs:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-amazon-ebs qemu:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-qemu azure-arm:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-azure-arm openstack:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-openstack amazon-ebsvolume:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-amazon-ebsvolume parallels-pvm:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-parallels-pvm docker:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-docker file:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-file vmware-vmx:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-vmware-vmx digitalocean:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-digitalocean profitbricks:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-profitbricks vmware-iso:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-vmware-iso amazon-ebssurrogate:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-amazon-ebssurrogate virtualbox-ovf:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-virtualbox-ovf amazon-chroot:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-amazon-chroot oneandone:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-oneandone null:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-null googlecompute:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-googlecompute triton:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-triton cloudstack:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-cloudstack hyperv-iso:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-hyperv-iso amazon-ebs-mock:MY_GOPATH/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock amazon-instance:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-amazon-instance parallels-iso:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-parallels-iso virtualbox-iso:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-virtualbox-iso] PostProcessors:map[artifice:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-artifice atlas:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-atlas checksum:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-checksum googlecompute-export:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-googlecompute-export docker-tag:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-docker-tag vsphere:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-vsphere vagrant-cloud:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-vagrant-cloud amazon-import:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-amazon-import shell-local:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-shell-local docker-import:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-docker-import compress:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-compress manifest:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-manifest docker-push:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-docker-push docker-save:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-docker-save vagrant:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-vagrant] Provisioners:map[powershell:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-powershell puppet-server:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-puppet-server salt-masterless:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-salt-masterless windows-shell:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-windows-shell file:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-file shell-local:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-shell-local chef-client:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-chef-client chef-solo:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-chef-solo converge:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-converge windows-restart:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-windows-restart puppet-masterless:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-puppet-masterless shell:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-shell ansible:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-ansible ansible-local:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-ansible-local]}
Setting cache directory: MY_GOPATH/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer_cache
Detected home directory from env var: /home/horgix
Loading builder: amazon-ebs-mock
Creating plugin client for path: MY_GOPATH/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock
Starting plugin: MY_GOPATH/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock []string{"MY_GOPATH/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock"}
Waiting for RPC address for: MY_GOPATH/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock
packer-builder-amazon-ebs-mock: Starting...
packer-builder-amazon-ebs-mock: Plugin minimum port: 10000
packer-builder-amazon-ebs-mock: Plugin maximum port: 25000
packer-builder-amazon-ebs-mock: Plugin address: unix /tmp/packer-plugin481514017
packer-builder-amazon-ebs-mock: Waiting for connection...
packer-builder-amazon-ebs-mock: Serving a plugin connection...
Build debug mode: false
Force build: false
On error:
Preparing build: amazon-ebs-mock
packer-builder-amazon-ebs-mock: Hello I'm a custom builder and this is the Prepare step
packer-builder-amazon-ebs-mock: Initializing random generator...
Waiting on builds to complete...
Starting build run: amazon-ebs-mock
Running builder: amazon-ebs-mock
packer-builder-amazon-ebs-mock: Hello I'm a custom builder
ui: ==> amazon-ebs-mock: I'm doing nothing...
==> amazon-ebs-mock: I'm doing nothing...
ui: ==> amazon-ebs-mock: Nothing done with success!
==> amazon-ebs-mock: Nothing done with success!
ui: ==> amazon-ebs-mock: Generating mock Artifact...
==> amazon-ebs-mock: Generating mock Artifact...
ui: ==> amazon-ebs-mock: Generated mock Artifact with success :)
==> amazon-ebs-mock: Generated mock Artifact with success :)
ui: Build 'amazon-ebs-mock' finished.
Builds completed. Waiting on interrupt barrier...
ui:
==> Builds finished. The artifacts of successful builds are:
machine readable: amazon-ebs-mock,artifact-count []string{"1"}
Build 'amazon-ebs-mock' finished.

==> Builds finished. The artifacts of successful builds are:
machine readable: amazon-ebs-mock,artifact []string{"0", "builder-id", ""}
machine readable: amazon-ebs-mock,artifact []string{"0", "id", "ap-northeast-1:ami-12345678,ap-northeast-2:ami-12345678,ap-southeast-1:ami-12345678,ap-southeast-2:ami-12345678,ca-central-1:ami-12345678,eu-west-1:ami-12345678,eu-west-2:ami-12345678,us-west-2:ami-12345678"}
machine readable: amazon-ebs-mock,artifact []string{"0", "string", "AMIs were created:\n\nap-northeast-1: ami-12345678\nap-northeast-2: ami-12345678\nap-southeast-1: ami-12345678\nap-southeast-2: ami-12345678\nca-central-1: ami-12345678\neu-west-1: ami-12345678\neu-west-2: ami-12345678\nus-west-2: ami-12345678"}
machine readable: amazon-ebs-mock,artifact []string{"0", "files-count", "0"}
machine readable: amazon-ebs-mock,artifact []string{"0", "end"}
ui: --> amazon-ebs-mock: AMIs were created:

ap-northeast-1: ami-12345678
ap-northeast-2: ami-12345678
ap-southeast-1: ami-12345678
ap-southeast-2: ami-12345678
ca-central-1: ami-12345678
eu-west-1: ami-12345678
eu-west-2: ami-12345678
us-west-2: ami-12345678
waiting for all plugin processes to complete...
--> amazon-ebs-mock: AMIs were created:

ap-northeast-1: ami-12345678
ap-northeast-2: ami-12345678
ap-southeast-1: ami-12345678
ap-southeast-2: ami-12345678
ca-central-1: ami-12345678
eu-west-1: ami-12345678
eu-west-2: ami-12345678
us-west-2: ami-12345678
MY_GOPATH/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock: plugin process exited
```
