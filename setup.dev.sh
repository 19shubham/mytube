#!/bin/bash
echo "!!!Start!!!"
echo "Creating image from Dockerfile"
docker build -f Dockerfile -t mytube . --build-arg env=dev
echo "!!!End!!!"