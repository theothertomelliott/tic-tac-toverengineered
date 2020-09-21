#!/bin/bash

# Build repository protos
protoc -I pkg/game/rpcrepository/ pkg/game/rpcrepository/repo.proto --go_out=plugins=grpc:pkg/game/rpcrepository --go_opt=module=github.com/theothertomelliott/tic-tac-toeverengineered/pkg/game/rpcrepository

# Build turn protos
protoc -I pkg/turn/rpcturn/ pkg/turn/rpcturn/current.proto --go_out=plugins=grpc:pkg/turn/rpcturn --go_opt=module=github.com/theothertomelliott/tic-tac-toeverengineered/pkg/turn/rpcturn