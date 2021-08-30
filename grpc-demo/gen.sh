#!/bin/bash


python -m grpc_tools.protoc -I./ --python_out=./ --grpc_python_out=./ item.proto
python -m grpc_tools.protoc -I./ --python_out=./ --grpc_python_out=./ order.proto
