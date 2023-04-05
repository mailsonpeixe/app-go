#!/bin/bash

NAMESPACE=lcd-25023
kubectl create namespace ${NAMESPACE}

echo "Create deployment"
kubectl apply -f kubernetes/deployment.yaml -n ${NAMESPACE}
echo "Create service"
kubectl apply -f kubernetes/service.yaml -n ${NAMESPACE}

