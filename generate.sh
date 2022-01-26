#!/bin/bash

protoc --go_out=gen --go_opt=paths=source_relative --go-grpc_out==plugins=grpc:gen --go-grpc_opt=paths=source_relative proto/greet.proto

