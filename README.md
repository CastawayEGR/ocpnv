ocpnv
=========
[![Go CI](https://github.com/CastawayEGR/ocpnv/actions/workflows/ci.yml/badge.svg)](https://github.com/CastawayEGR/ocpnv/actions/workflows/ci.yml)
[![Go Release](https://github.com/CastawayEGR/ocpnv/actions/workflows/release.yml/badge.svg)](https://github.com/CastawayEGR/ocpnv/actions/workflows/release.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![GitHub repo size in bytes](https://img.shields.io/github/repo-size/CastawayEGR/ocpnv.svg?logoColor=brightgreen)](https://github.com/CastawayEGR/ocpnv)
[![GitHub last commit](https://img.shields.io/github/last-commit/CastawayEGR/ocpnv.svg?logoColor=brightgreen)](https://github.com/CastawayEGR/ocpnv)

Generate OCP complaint MachineConfigs from RHEL Entitlements Certificates

## Overview

`ocpnv` pronounced OCP-N-V (OCP Envy) is a very basic program to convert a downloaded zip file from access.redhat.com
into MachineConfig yaml for OpenShift to enable RHCOS access to additional repos.

## Installation

`ocpnv` is available from the [project's releases page](https://github.com/castawayegr/ocpnv/releases).

## Usage

```bash
# Create MachineConfig yaml config file
$ ocpnv -f downloaded_entitlements.zip > template.yaml

# Import into your OpenShift cluster
$ oc create -f template.yaml
```
```bash
# Create and import MachineConfig in one shot
$ ocpnv -f downloaded_entitlements.zip | oc create -f -
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
