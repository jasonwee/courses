#!/usr/bin/python

import sys
import grpc
import demo_pb2
import demo_pb2_grpc
from grpc.health.v1 import health_pb2
from grpc.health.v1 import health_pb2_grpc

from logger import getJSONLogger
logger = getJSONLogger('adservice-v2-server')

if __name__ == "__main__":

    # TODO
    # set up server stub
    # ensure the service is listening to port 9556
    channel = grpc.insecure_channel("localhost:9556")

    # TODO
    # form a request with the required input
    adRequest = demo_pb2.AdRequest(
        context_keys="1234"
    )

    # TODO
    # make a call to server and return a response
    stub = demo_pb2_grpc.AdServiceV2Stub(channel)
    response = stub.GetAdsV2(adRequest)

    # Uncomment to log the response from the Server
    logger.info(response)
