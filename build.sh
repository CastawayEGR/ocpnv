#!/bin/bash
podman build -t gobuild .
podman container create --name temp gobuild
podman container cp temp:/ocpnv-amd64-linux .
podman container cp temp:/ocpnv-amd64-darwin .
podman container cp temp:/ocpnv-amd64.exe .
podman rm temp
