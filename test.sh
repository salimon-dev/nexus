#!/bin/bash

echo "Running Go tests..."

go test ./helpers -coverprofile=helpers-cover.out
go tool cover -html=helpers-cover.out -o helpers-cover.html