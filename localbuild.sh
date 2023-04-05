#!/bin/bash

docker build -t app-go .

docker run -p 3000:3000 app-go

#URL access http://localhost:3000