#!/bin/bash

# Setteo variables
DOCKER="sudo docker"
IMAGE_NAME="golang_present:latest"
CONTAINER_NAME="presentacion"

# Me aseguro de no tener un container que se llame $CONTAINER_NAME
$DOCKER stop $CONTAINER_NAME 
$DOCKER rm $CONTAINER_NAME 

#Buildeo la imagen
$DOCKER build -t $IMAGE_NAME .

#Inicio el container
$DOCKER run -d --name $CONTAINER_NAME -p 3999:3999 $IMAGE_NAME

#Copio la presentacion dentro del container
$DOCKER cp presentacion.slide $CONTAINER_NAME:/go/src/app
$DOCKER cp boring3.go $CONTAINER_NAME:/go/src/app