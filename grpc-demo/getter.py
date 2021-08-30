import grpc
import order_pb2
import order_pb2_grpc

print("sending sample payload...")

channel = grpc.insecure_channel("localhost:5005")
stub = order_pb2_grpc.OrderServiceStub(channel)

response = stub.Get(order_pb2.Empty())
print(response)
