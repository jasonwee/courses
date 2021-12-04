#!/bin/bash -eu

set -e

source py39_env/bin/activate
# TODO: Add the commands to generate the gRPC files
`python -m grpc_tools.protoc -I../../pb --python_out=./ --grpc_python_out=./ demo.proto`
