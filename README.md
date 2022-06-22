# Setup

## Necessary Setup

* Install [Docker-Desktop](https://www.docker.com/products/docker-desktop/) or [Rancher-Desktop](https://rancherdesktop.io/)

* Install [Go](https://go.dev/doc/install)

* Install [DotNet](https://docs.microsoft.com/en-us/dotnet/core/install/linux-ubuntu#2004)

* Install [Buf](https://docs.buf.build/installation)

* Install [Helm](https://helm.sh/docs/intro/install/)

* Install node (optionally, using [nvm](https://github.com/nvm-sh/nvm/blob/master/README.md#installing-and-updating))

* Install [yarn](https://classic.yarnpkg.com/lang/en/docs/install/)

* Install necessary tools:
```bash
# GCC
sudo apt install gcc -y

# Make
sudo apt install make -y

# WSL-U (WSL2 systems only) (https://wslutiliti.es/wslu/install.html)
sudo apt install ubuntu-wsl

# You'll need to add this to your bashrc (or profile, I'm not a cop) if you've installed WSL-U:
# export BROWSER=wslview

# Tilt
curl -fsSL https://raw.githubusercontent.com/tilt-dev/tilt/master/scripts/install.sh | bash
```

* Install protobuf plugins:
```bash
# Protobuf Plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/envoyproxy/protoc-gen-validate@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

## Optional Setup

* Install [Loft](https://loft.sh/docs/getting-started/install/cli)

* Install [Evans](https://github.com/ktr0731/evans#readme)

The following extensions are recommended for VS Code development:
* [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
* [Docker](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)
* [Kubernetes](https://marketplace.visualstudio.com/items?itemName=ms-kubernetes-tools.vscode-kubernetes-tools)
* [Tiltfile](https://marketplace.visualstudio.com/items?itemName=tilt-dev.Tiltfile)
* [vscode-proto3](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3)
* [YAML](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml)

If you've installed the vscode-proto3 extension, you may need to add the following to your local
settings.json ([ref](https://github.com/zxh0/vscode-proto3/issues/31#issuecomment-409175121)):
```json
{
  "protoc": {
    "options": [
      "--proto_path=api"
    ]
  }
}
```

# Architecture

[![](https://mermaid.ink/img/pako:eNp1kU1vwjAMhv9KlNMmgbj3MAlavm5osFPDIW1cGkESljibJsR_n9t0AiTWS6s8r5_Y9YXXTgHP-MHLc8t2hbCMnmkpxEcAz3ZeNo2uJ3s2HjNtNWqJwIh8EQzoQRptDwTf2Gwo7ZIBrGK1M0bS-w4W5Qwsti70vpMOCJah6yNzurPQoXZe3a5NZSFWqb9lvkkns8t0s752ksb5b0klt2Y8fEYI2Evz8t05s3X1EXCQFU8bXAzee-UD70V5OhlMiy58jhXN0T6Ely_lJlaTbaz2rym66svXtnFPJl-mTP4v6bs6ADJPElpDk-gqURpFWD7iBryRWtEyLx0QHFswIHhGn0r6o-DCXikXz4qWOFcanedZI08BRlxGdNsfW_MMfYS_UKEl_XgzpK6_kkSufA)](https://mermaid.live/edit#pako:eNp1kU1vwjAMhv9KlNMmgbj3MAlavm5osFPDIW1cGkESljibJsR_n9t0AiTWS6s8r5_Y9YXXTgHP-MHLc8t2hbCMnmkpxEcAz3ZeNo2uJ3s2HjNtNWqJwIh8EQzoQRptDwTf2Gwo7ZIBrGK1M0bS-w4W5Qwsti70vpMOCJah6yNzurPQoXZe3a5NZSFWqb9lvkkns8t0s752ksb5b0klt2Y8fEYI2Evz8t05s3X1EXCQFU8bXAzee-UD70V5OhlMiy58jhXN0T6Ely_lJlaTbaz2rym66svXtnFPJl-mTP4v6bs6ADJPElpDk-gqURpFWD7iBryRWtEyLx0QHFswIHhGn0r6o-DCXikXz4qWOFcanedZI08BRlxGdNsfW_MMfYS_UKEl_XgzpK6_kkSufA)

<details>
<summary>Mermaid.live Source</summary>
```
graph TD
    A[\User Traffic/] -- initiate server streaming --> B
    A -- send command --> B
    D[Benthos] -- listen to --> E[\Discord Traffic/]
    subgraph GCP
    B{API} -- forward streaming request --> C[RoomSocket]
    D -- send command --> F
    B -- forward command --> F[RoomCommand]
    F -- publish command --> G([Pub/Sub])
    H[RoomInfo] -- listen to --> G
    C -- listen to --> G
    B -- get room info --> H
    end
```
</details>
