from concurrent import futures
import time

import grpc

import users_pb2
import users_pb2_grpc
import json

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

data = json.load(open('users.json'))

class UsersService(users_pb2_grpc.UsersServicer):
    def RequestUser(self, request, context):
        filterResult = list(filter(lambda user: user['id'] == request.userId, data))
        userDic = filterResult[0]
        return users_pb2.UserResponse(name = userDic['name'], imageUrl = userDic['imgUrl'])

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    users_pb2_grpc.add_UsersServicer_to_server(UsersService(), server)
    server.add_insecure_port('[::]:22222')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
  serve()
