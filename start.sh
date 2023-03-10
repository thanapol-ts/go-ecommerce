#!/bin/sh

IMAGE_NAME="go-app"
CONTAINER_NAME="go-ecommerce"
PORT_NUMBER="4000"
ENV= "dev"

docker build -t $IMAGE_NAME .

docker run -e ENV=$ENV -d -p $PORT_NUMBER:$PORT_NUMBER -it $IMAGE_NAME

exit 1
