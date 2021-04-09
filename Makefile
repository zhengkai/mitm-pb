SHELL:=/bin/bash

default:
	./server/build/run-server.sh

stop:
	./server/build/stop-server.sh
