FROM ubuntu:latest

RUN apt update && apt install -y software-properties-common
RUN add-apt-repository ppa:deadsnakes/ppa -y
RUN apt update
RUN apt install -y python3.10
RUN python3.10 --version
