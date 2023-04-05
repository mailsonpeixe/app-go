#!/bin/bash
IMAGE=mailsonpeixe/app-go
echo "Build and push image ${IMAGE}"

docker build --push -t $IMAGE:1.0 .

