import event_coordinates_pb2
import event_coordinates_pb2_grpc
import grpc

# Simulates user device sending coordinates to gRPC

print("Sending coordinates...")

channel = grpc.insecure_channel("127.0.0.1:30003")
stub = event_coordinates_pb2_grpc.EventCoordinatesServiceStub(channel)

# Update this with desired payload
user_coordinates = event_coordinates_pb2.EventCoordinatesMessage(
    userId=1,
    latitude=-100,
    longitude=30
)

user_coordinates_2 = event_coordinates_pb2.EventCoordinatesMessage(
    userId=9,
    latitude=-100,
    longitude=30
)

response_1 = stub.Create(user_coordinates)
response_2 = stub.Create(user_coordinates_2)


print("Coordinates sent...")
print(user_coordinates, user_coordinates_2)
