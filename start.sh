#!/bin/sh

IMAGE_NAME="go-app"
CONTAINER_NAME="go-ecommerce"
PORT_NUMBER="4000"
ENV= "dev"

docker images

docker build -t $IMAGE_NAME .

docker ps -all

docker run -e ENV=$ENV -d -p $PORT_NUMBER:$PORT_NUMBER --name $CONTAINER_NAME -it $IMAGE_NAME
