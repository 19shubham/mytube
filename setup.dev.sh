#!/bin/bash
echo "!!!Start!!!"
echo "Creating image from Dockerfile"
docker build -t mytube .
echo "!!!End!!!"