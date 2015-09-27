FROM debian:stretch

WORKDIR /opt/evolution-master
RUN apt-get update
RUN apt-get install -y gcc python-dev libgit2-dev

COPY ./ /opt/evolution-master
