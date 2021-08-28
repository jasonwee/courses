## Generating gRPC files
`python3.9 -m venv py39_env`
`source py39_env/bin/activate`
`pip3 install --upgrade pip`
`pip install grpcio-tools grpcio`

`python -m grpc_tools.protoc -I./ --python_out=./ --grpc_python_out=./ item.proto`
