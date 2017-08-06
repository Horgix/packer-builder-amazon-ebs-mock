**This Packer plugin is a Builder that mocks the [official Amazon EBS
builder](https://www.packer.io/docs/builders/amazon-ebs.html).**

Why do you want that? No idea, for testing purpose? At least this was my case,
and it was built to help me test my [Packer post-processor that sends Slack
notifications](https://github.com/Horgix/packer-post-processor-slack-notifications);
I obviously don't want to wait for a real AMI to build everytime each time I
want to run a test...

# How to use it

Basically refer to the [Packer documentation on building
images](https://www.packer.io/intro/getting-started/build-image.html) about how
to call it in your `packer.json`, since it's working as any builder. To learn
how to **install** this plugin, refer to the [Packer documentation on
installing
plugins](https://www.packer.io/docs/extending/plugins.html#installing-plugins)

# How to build it

```shell
# Get dependencies
glide install
# Build the binary
make
```

# Example

## Most simple Packer JSON input (available in `tests/simplest_packer.json`):

```json
{
  "builders": [
    {
      "type":           "amazon-ebs-mock"
    }
  ]
}
```

## Output of such a build


`packer build ./tests/simplest_packer.json`

```
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

## Want more details? Enable Packer logging!

See at the end of this README, it's huge :)

# Misc informations

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

## Samples of `manifest.json` produced by the `manifest` post-processor with this builder


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

# Details of a basic build


`PACKER_LOG=true packer build ./tests/simplest_packer.json`

TODO : Remove the color codes from that...

```
2017/08/04 20:01:56 [INFO] Packer version: 0.12.3
2017/08/04 20:01:56 Packer Target OS/Arch: linux amd64
2017/08/04 20:01:56 Built with Go Version: go1.8
2017/08/04 20:01:56 Detected home directory from env var: /home/horgix
2017/08/04 20:01:56 [DEBUG] Discovered plugin: amazon-ebs-mock = /home/horgix/work/godev/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock
2017/08/04 20:01:56 Using internal plugin for amazon-ebssurrogate
2017/08/04 20:01:56 Using internal plugin for profitbricks
2017/08/04 20:01:56 Using internal plugin for triton
2017/08/04 20:01:56 Using internal plugin for virtualbox-ovf
2017/08/04 20:01:56 Using internal plugin for vmware-vmx
2017/08/04 20:01:56 Using internal plugin for null
2017/08/04 20:01:56 Using internal plugin for oneandone
2017/08/04 20:01:56 Using internal plugin for openstack
2017/08/04 20:01:56 Using internal plugin for qemu
2017/08/04 20:01:56 Using internal plugin for vmware-iso
2017/08/04 20:01:56 Using internal plugin for amazon-chroot
2017/08/04 20:01:56 Using internal plugin for amazon-ebsvolume
2017/08/04 20:01:56 Using internal plugin for azure-arm
2017/08/04 20:01:56 Using internal plugin for googlecompute
2017/08/04 20:01:56 Using internal plugin for hyperv-iso
2017/08/04 20:01:56 Using internal plugin for parallels-pvm
2017/08/04 20:01:56 Using internal plugin for virtualbox-iso
2017/08/04 20:01:56 Using internal plugin for amazon-ebs
2017/08/04 20:01:56 Using internal plugin for amazon-instance
2017/08/04 20:01:56 Using internal plugin for cloudstack
2017/08/04 20:01:56 Using internal plugin for digitalocean
2017/08/04 20:01:56 Using internal plugin for docker
2017/08/04 20:01:56 Using internal plugin for file
2017/08/04 20:01:56 Using internal plugin for parallels-iso
2017/08/04 20:01:56 Using internal plugin for puppet-masterless
2017/08/04 20:01:56 Using internal plugin for ansible-local
2017/08/04 20:01:56 Using internal plugin for powershell
2017/08/04 20:01:56 Using internal plugin for puppet-server
2017/08/04 20:01:56 Using internal plugin for windows-restart
2017/08/04 20:01:56 Using internal plugin for shell
2017/08/04 20:01:56 Using internal plugin for shell-local
2017/08/04 20:01:56 Using internal plugin for file
2017/08/04 20:01:56 Using internal plugin for salt-masterless
2017/08/04 20:01:56 Using internal plugin for windows-shell
2017/08/04 20:01:56 Using internal plugin for ansible
2017/08/04 20:01:56 Using internal plugin for chef-client
2017/08/04 20:01:56 Using internal plugin for chef-solo
2017/08/04 20:01:56 Using internal plugin for converge
2017/08/04 20:01:56 Using internal plugin for atlas
2017/08/04 20:01:56 Using internal plugin for docker-save
2017/08/04 20:01:56 Using internal plugin for docker-tag
2017/08/04 20:01:56 Using internal plugin for googlecompute-export
2017/08/04 20:01:56 Using internal plugin for vagrant-cloud
2017/08/04 20:01:56 Using internal plugin for vsphere
2017/08/04 20:01:56 Using internal plugin for compress
2017/08/04 20:01:56 Using internal plugin for checksum
2017/08/04 20:01:56 Using internal plugin for docker-import
2017/08/04 20:01:56 Using internal plugin for amazon-import
2017/08/04 20:01:56 Using internal plugin for docker-push
2017/08/04 20:01:56 Using internal plugin for manifest
2017/08/04 20:01:56 Using internal plugin for shell-local
2017/08/04 20:01:56 Using internal plugin for vagrant
2017/08/04 20:01:56 Using internal plugin for artifice
2017/08/04 20:01:56 Detected home directory from env var: /home/horgix
2017/08/04 20:01:56 Attempting to open config file: /home/horgix/.packerconfig
2017/08/04 20:01:56 [WARN] Config file doesn't exist: /home/horgix/.packerconfig
2017/08/04 20:01:56 Packer config: &{DisableCheckpoint:false DisableCheckpointSignature:false PluginMinPort:10000 PluginMaxPort:25000 Builders:map[qemu:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-qemu amazon-chroot:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-amazon-chroot virtualbox-iso:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-virtualbox-iso amazon-ebs:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-amazon-ebs digitalocean:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-digitalocean triton:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-triton oneandone:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-oneandone openstack:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-openstack docker:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-docker cloudstack:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-cloudstack file:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-file virtualbox-ovf:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-virtualbox-ovf parallels-pvm:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-parallels-pvm amazon-instance:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-amazon-instance amazon-ebsvolume:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-amazon-ebsvolume hyperv-iso:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-hyperv-iso amazon-ebssurrogate:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-amazon-ebssurrogate profitbricks:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-profitbricks vmware-vmx:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-vmware-vmx azure-arm:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-azure-arm googlecompute:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-googlecompute parallels-iso:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-parallels-iso amazon-ebs-mock:/home/horgix/work/godev/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock null:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-null vmware-iso:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-builder-vmware-iso] PostProcessors:map[artifice:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-artifice docker-save:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-docker-save docker-tag:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-docker-tag checksum:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-checksum manifest:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-manifest vagrant:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-vagrant atlas:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-atlas googlecompute-export:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-googlecompute-export vagrant-cloud:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-vagrant-cloud amazon-import:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-amazon-import docker-push:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-docker-push compress:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-compress vsphere:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-vsphere docker-import:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-docker-import shell-local:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-post-processor-shell-local] Provisioners:map[powershell:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-powershell shell-local:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-shell-local file:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-file salt-masterless:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-salt-masterless converge:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-converge windows-restart:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-windows-restart puppet-masterless:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-puppet-masterless windows-shell:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-windows-shell chef-solo:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-chef-solo ansible-local:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-ansible-local puppet-server:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-puppet-server shell:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-shell ansible:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-ansible chef-client:/usr/bin/packer-io-PACKERSPACE-plugin-PACKERSPACE-packer-provisioner-chef-client]}
2017/08/04 20:01:56 Setting cache directory: /home/horgix/work/godev/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer_cache
2017/08/04 20:01:56 Detected home directory from env var: /home/horgix
2017/08/04 20:01:56 Loading builder: amazon-ebs-mock
2017/08/04 20:01:56 Creating plugin client for path: /home/horgix/work/godev/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock
2017/08/04 20:01:56 Starting plugin: /home/horgix/work/godev/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock []string{"/home/horgix/work/godev/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock"}
2017/08/04 20:01:56 Waiting for RPC address for: /home/horgix/work/godev/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock
2017/08/04 20:01:56 packer-builder-amazon-ebs-mock: 2017/08/04 20:01:56 Starting...
2017/08/04 20:01:56 packer-builder-amazon-ebs-mock: 2017/08/04 20:01:56 Plugin minimum port: 10000
2017/08/04 20:01:56 packer-builder-amazon-ebs-mock: 2017/08/04 20:01:56 Plugin maximum port: 25000
2017/08/04 20:01:56 packer-builder-amazon-ebs-mock: 2017/08/04 20:01:56 Plugin address: unix /tmp/packer-plugin264942503
2017/08/04 20:01:56 packer-builder-amazon-ebs-mock: 2017/08/04 20:01:56 Waiting for connection...
2017/08/04 20:01:56 packer-builder-amazon-ebs-mock: 2017/08/04 20:01:56 Serving a plugin connection...
2017/08/04 20:01:56 ui: [1;32mamazon-ebs-mock output will be in this color.[0m
[1;32mamazon-ebs-mock output will be in this color.[0m
2017/08/04 20:01:56 ui: 

2017/08/04 20:01:56 Build debug mode: false
2017/08/04 20:01:56 Force build: false
2017/08/04 20:01:56 On error: 
2017/08/04 20:01:56 Preparing build: amazon-ebs-mock
2017/08/04 20:01:56 packer-builder-amazon-ebs-mock: 2017/08/04 20:01:56 Hello I'm a custom builder and this is the Prepare step
2017/08/04 20:01:56 packer-builder-amazon-ebs-mock: 2017/08/04 20:01:56 Initializing random generator...
2017/08/04 20:01:56 Waiting on builds to complete...
2017/08/04 20:01:56 Starting build run: amazon-ebs-mock
2017/08/04 20:01:56 Running builder: amazon-ebs-mock
2017/08/04 20:01:56 packer-builder-amazon-ebs-mock: 2017/08/04 20:01:56 Hello I'm a custom builder
2017/08/04 20:01:56 ui: [1;32m==> amazon-ebs-mock: I'm doing nothing...[0m
[1;32m==> amazon-ebs-mock: I'm doing nothing...[0m
2017/08/04 20:01:56 ui: [1;32m==> amazon-ebs-mock: Nothing done with success![0m
[1;32m==> amazon-ebs-mock: Nothing done with success![0m
2017/08/04 20:01:56 ui: [1;32m==> amazon-ebs-mock: Generating mock Artifact...[0m
[1;32m==> amazon-ebs-mock: Generating mock Artifact...[0m
2017/08/04 20:01:56 ui: [1;32m==> amazon-ebs-mock: Generated mock Artifact with success :)[0m
[1;32m==> amazon-ebs-mock: Generated mock Artifact with success :)[0m
2017/08/04 20:01:56 ui: [1;32mBuild 'amazon-ebs-mock' finished.[0m
2017/08/04 20:01:56 Builds completed. Waiting on interrupt barrier...
2017/08/04 20:01:56 ui: 
==> Builds finished. The artifacts of successful builds are:
2017/08/04 20:01:56 machine readable: amazon-ebs-mock,artifact-count []string{"1"}
[1;32mBuild 'amazon-ebs-mock' finished.[0m

==> Builds finished. The artifacts of successful builds are:
2017/08/04 20:01:56 machine readable: amazon-ebs-mock,artifact []string{"0", "builder-id", ""}
2017/08/04 20:01:56 machine readable: amazon-ebs-mock,artifact []string{"0", "id", "ap-south-1:ami-12345678,ap-southeast-1:ami-12345678,ap-southeast-2:ami-12345678,eu-west-1:ami-12345678,eu-west-2:ami-12345678,us-east-2:ami-12345678,us-west-2:ami-12345678"}
2017/08/04 20:01:56 machine readable: amazon-ebs-mock,artifact []string{"0", "string", "AMIs were created:\n\nap-south-1: ami-12345678\nap-southeast-1: ami-12345678\nap-southeast-2: ami-12345678\neu-west-1: ami-12345678\neu-west-2: ami-12345678\nus-east-2: ami-12345678\nus-west-2: ami-12345678"}
2017/08/04 20:01:56 machine readable: amazon-ebs-mock,artifact []string{"0", "files-count", "0"}
2017/08/04 20:01:56 machine readable: amazon-ebs-mock,artifact []string{"0", "end"}
2017/08/04 20:01:56 ui: --> amazon-ebs-mock: AMIs were created:

ap-south-1: ami-12345678
ap-southeast-1: ami-12345678
ap-southeast-2: ami-12345678
eu-west-1: ami-12345678
eu-west-2: ami-12345678
us-east-2: ami-12345678
us-west-2: ami-12345678
2017/08/04 20:01:56 waiting for all plugin processes to complete...
--> amazon-ebs-mock: AMIs were created:

ap-south-1: ami-12345678
ap-southeast-1: ami-12345678
ap-southeast-2: ami-12345678
eu-west-1: ami-12345678
eu-west-2: ami-12345678
us-east-2: ami-12345678
us-west-2: ami-12345678
2017/08/04 20:01:56 /home/horgix/work/godev/src/github.com/horgix/packer-builder-amazon-ebs-mock/packer-builder-amazon-ebs-mock: plugin process exited
