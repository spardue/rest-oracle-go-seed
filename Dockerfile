# syntax=docker/dockerfile:1

FROM golang:1.17

RUN go get github.com/gin-gonic/gin


RUN export LD_LIBRARY_PATH=/usr/lib/oracle/19.13/client64/lib
RUN export DPI_DEBUG_LEVEL=16

