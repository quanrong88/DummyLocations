FROM grpc/python:1.4
ADD . /code
WORKDIR /code
RUN pip install --upgrade gcloud
RUN pip install --upgrade google-api-python-client
EXPOSE 22222
CMD ["python", "server.py"]
