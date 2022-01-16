# syntax=docker/dockerfile:1
FROM ubuntu:latest

ENV LD_LIBRARY_PATH=/lib

USER root

WORKDIR /root

# Install necessary command line utils 
RUN apt update
RUN DEBIAN_FRONTEND=noninteractive TZ=Etc/UTC apt-get -y install tzdata
RUN apt -y install wget
RUN apt -y install unzip
RUN apt -y install alien

# Install Go 1.17.6
RUN wget https://go.dev/dl/go1.17.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf ./go*.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin

# Install Oracle Drivers
RUN wget https://download.oracle.com/otn_software/linux/instantclient/1913000/oracle-instantclient19.13-basic-19.13.0.0.0-2.x86_64.rpm
RUN alien -i oracle-instantclient*.x86_64.rpm
RUN apt install -y libaio1 libaio-dev

# Setup runtime directory
RUN mkdir -p /app
COPY ./ /app
WORKDIR /app
RUN /usr/local/go/bin/go mod download

# Normal run time intended for production. 
EXPOSE 8080
CMD LD_LIBRARY_PATH=/usr/lib/oracle/19.13/client64/lib /usr/local/go/bin/go run .
#CMD LD_LIBRARY_PATH=/usr/lib/oracle/19.13/client64/lib:DPI_DEBUG_LEVEL=16 /usr/local/go/bin run .
