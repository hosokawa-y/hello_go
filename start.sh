#!/bin/bash

# generate swagger documents
echo "generate swagger documents..."
swag init -g app/controllers/server.go
echo "go run..."
go run main.go
