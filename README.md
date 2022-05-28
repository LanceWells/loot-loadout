# Setup

Install Docker-Desktop (or Rancher-Desktop)

Install Tilt:
  ```
  curl -fsSL https://raw.githubusercontent.com/tilt-dev/tilt/master/scripts/install.sh | bash
  ```

You may need to add the following to VSCode settings:

[Reference](https://earthly.dev/blog/golang-monorepo/#:~:text=of%20the%20article.-,Monorepo%20Layout%20in%20Go,and%20a%20single%20shared%20Library.)

```
"gopls": {
  "experimentalWorkspaceModule": true
}
```

Install protobuf:
```
sudo apt install protobuf-compiler
sudo apt install golang-goprotobuf-dev
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
