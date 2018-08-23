#!/bin/bash

if [[ "$#" -ne 2 ]]; then
	echo "To Test Local Relay: test_server -r <port>"
	echo "To Test Local Exit:  test_server -e <port>"
elif [[ "$1" = "-r" ]]; then
	curl -X POST -H "Content-Type: application/json" -d @relay_request.json localhost:${2}
elif [[ "$1" = "-e" ]]; then
	curl -X POST -H "Content-Type: application/json" -d @exit_request.json localhost:${2}
fi
