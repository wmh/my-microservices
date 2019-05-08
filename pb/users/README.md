
## go

```
protoc -I . --go_out=plugins=grpc:. ./users.proto
```

## php

```
protoc --proto_path=. \
  --php_out=./php \
  --grpc_out=./php \
  --plugin=protoc-gen-grpc=/usr/local/bin/grpc_php_plugin \
  ./users.proto
```