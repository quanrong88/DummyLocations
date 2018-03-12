from concurrent import futures
import time

import grpc

import locations_pb2
import locations_pb2_grpc
import comments_pb2
import comments_pb2_grpc

import json

_ONE_DAY_IN_SECONDS = 60 * 60 * 24
data = json.load(open('locations.json'))

class LocationsService(locations_pb2_grpc.LocationsServicer):
    def RequestLocations(self, request, context):
        response = locations_pb2.LocationsResponse()
        for item in data:
            location = response.locations.add()
            location.id = item['id']
            location.title = item['title']
            location.subtitle = item['subtitle']
            location.lat = item['lat']
            location.long = item['long']
        return response
    def RequestLocationDetail(self, request, context):
        contentId = request.locationId
        filterResult = list(filter(lambda location: location['id'] == contentId, data))
        locationDic = filterResult[0]
        channel = grpc.insecure_channel('comments-module:22222')
        stub = comments_pb2_grpc.CommentsStub(channel)
        response = stub.RequestComments(comments_pb2.CommentsRequest(contentId = contentId))
        # locationRequest = locations_pb2.Location()
        # locationRequest.id = locationDic['id']
        # locationRequest.title = locationDic['title']
        # locationRequest.subtitle = locationDic['subtitle']
        # locationRequest.lat = locationDic['lat']
        # locationRequest.long = locationDic['long']
        detailResoponse = locations_pb2.LocationDetailResponse(location = locations_pb2.Location(id = locationDic['id'], title = locationDic['title'], subtitle = locationDic['subtitle'], lat = locationDic['lat'], long = locationDic['long']))
        for item in response.comments:
            comment = detailResoponse.comment.add()
            comment.content = item.content
            comment.name = item.name
            comment.imageUrl = item.imageUrl
        return detailResoponse

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    locations_pb2_grpc.add_LocationsServicer_to_server(LocationsService(), server)
    server.add_insecure_port('[::]:22222')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
  serve()
