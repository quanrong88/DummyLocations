from __future__ import print_function

import grpc
import locations_pb2
import locations_pb2_grpc
from pprint import pprint

def run():
    channel = grpc.insecure_channel('locations-module:22222')
    stub = locations_pb2_grpc.LocationsStub(channel)
    response = stub.RequestLocationDetail(locations_pb2.LocationDetailRequest(locationId = 1))
    pprint(response)

if __name__ == '__main__':
  run()
