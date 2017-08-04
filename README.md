This Packer plugin is a **Builder** that mocks the [official Amazon EBS
builder](https://www.packer.io/docs/builders/amazon.html TODO check link). Why
do you want that? No idea, for testing purpose? At least this was my case, and
it was built to help me test my [Packer post-processor that sends Slack
notifications](TODO LINK) (I obviously don't want to wait for a real AMI to
build everytime each time I want to run a test...)

# How to use it

Basically refer to the TODO link Packer documentation

-

# How to build it

```shell
# Get dependencies
glide install
# Build the binary
make
```

# Example

## Most simple Packer JSON input (available in `tests/simplest_packer.json`):

```
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

```
