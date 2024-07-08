#!/bin/sh
docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama
docker cp .\Modelfile ollama:/
docker exec -it ollama ollama run create test -f ./Modelfile