```
python3.9 -m venv py39_env
source py39_env/bin/activate
pip3 install --upgrade pip
pip install grpcio-tools grpcio
```

```
bash genproto.sh
```

```
source py39_env/bin/activate
mkdir -p grpc/health/v1
bash genhealth.sh
```
