#!/bin/bash

export KAFKA_URL=my-release-kafka-0.my-release-kafka-headless.default.svc.cluster.local:9092
export KAFKA_TOPIC=test
python grpc_server_startup.py
