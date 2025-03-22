FROM ubuntu:latest
LABEL authors="arsal"

ENTRYPOINT ["top", "-b"]