# This docker file builds an image that runs curl
# $ docker build -t meaningfulname .
# $ docker images
# $ docker run -it meaningfulname /bin/bash
# # curl --head www.google.com
FROM ubuntu:latest
RUN apt-get -y update && apt-get install -y curl
