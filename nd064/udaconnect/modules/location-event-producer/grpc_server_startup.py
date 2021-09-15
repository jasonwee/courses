import event_coordinates_pb2
import event_coordinates_pb2_grpc
import grpc
import json
import logging
import os

from concurrent import futures
from kafka import KafkaProducer

kafka_url = os.environ["KAFKA_URL"]
kafka_topic = os.environ["KAFKA_TOPIC"]

logging.info('connecting to kafka ', kafka_url)
print('connecting to kafka ', kafka_url)
logging.info('connecting to kafka topic ', kafka_topic)
print('connecting to kafka topic ', kafka_topic)

producer = KafkaProducer(bootstrap_servers=kafka_url)

class EventCoordinatesServicer(event_coordinates_pb2_grpc.EventCoordinatesServiceServicer):

    def Create(self, request, context):
        request_value = {
            'userId': int(request.userId),
            'latitude': int(request.latitude),
            'longitude': int(request.longitude)
        }

        logging.info('processing entity ', request_value)

        user_encode_data = json.dumps(request_value, indent=2).encode('utf-8')
        producer.send(kafka_topic, user_encode_data)
        return event_coordinates_pb2.EventCoordinatesMessage(**request_value)


server = grpc.server(futures.ThreadPoolExecutor(max_workers=2))
event_coordinates_pb2_grpc.add_EventCoordinatesServiceServicer_to_server(EventCoordinatesServicer(), server)

logging.info('starting on port 5005')
server.add_insecure_port('[::]:5005')
server.start()
server.wait_for_termination()
