
### setup python env
```
python3.9 -m venv py39_env
source py39_env/bin/activate
pip3 install --upgrade pip
pip install -r requirements.txt

```

### compile
```
python -m grpc_tools.protoc -I./ --python_out=./ --grpc_python_out=./ event_coordinates.proto
```

### run locally
```
$ run-local.sh
```


## Docker commands used to build the application 
docker build --no-cache -t jasonwee/udaconnect-location-event-producer .


## Docker commands used to run the application
docker run -e KAFKA_URL=my-release-0.kafka-release-headless.default.svc.cluster.local:9092 -e KAFKA_TOPIC=test -d -p 5000:5000 --name udaconnect-location-event-producer --rm jasonwee/udaconnect-location-event-producer


## Docker commands used to get the application logs
docker logs -f udaconnect-location-event-producer


## push to dockerhub
docker push jasonwee/udaconnect-location-event-producer
