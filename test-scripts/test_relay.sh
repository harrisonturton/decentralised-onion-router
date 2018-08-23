#!/bin/bash

if [[ "$#" -ne 1 ]]; then
	echo "Usage: test_server <-r/-e>"
elif [[ "$1" = "-r" ]]; then
	curl -X POST -H "Content-Type: application/json" -d @relay_request.json localhost:3333
elif [[ "$1" = "-e" ]]; then
	curl -X POST -H "Content-Type: application/json" -d @exit_request.json localhost:3333
fi
