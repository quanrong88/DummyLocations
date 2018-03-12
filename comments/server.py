from concurrent import futures
import time

import grpc

import comments_pb2
import comments_pb2_grpc
import users_pb2
import users_pb2_grpc

import json

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

data = json.load(open('comments.json'))

class CommentsService(comments_pb2_grpc.CommentsServicer):
    def RequestComments(self, request, context):
        contentId = request.contentId
        comments = list(filter(lambda comment: comment['contentId'] == contentId, data))
        response = comments_pb2.CommentsResponse()
        channel = grpc.insecure_channel('users-module:22222')
        stub = users_pb2_grpc.UsersStub(channel)
        for item in comments:
            comment = response.comments.add()
            comment.content = item['comment']
            responseUser = stub.RequestUser(users_pb2.UserRequest(userId = item['userId']))
            comment.name = responseUser.name
            comment.imageUrl = responseUser.imageUrl
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    comments_pb2_grpc.add_CommentsServicer_to_server(CommentsService(), server)
    server.add_insecure_port('[::]:22222')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
  serve()
