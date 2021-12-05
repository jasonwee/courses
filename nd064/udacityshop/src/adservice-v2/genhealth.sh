#!/bin/bash -eu

set -e

# TODO: Add the commands to generate the gRPC files
#python -m grpc_tools.protoc -I../../pb/grpc/health/v1/ --python_out=./grpc/health/v1/ --grpc_python_out=./grpc/health/v1/ health.proto
source py39_env/bin/activate
python -m grpc_tools.protoc -I ../../pb/ --python_out=./grpc/health/v1 --grpc_python_out=./grpc/health/v1 grpc/health/v1/health.proto
