FROM debian:stretch

WORKDIR /opt/evolution-master
RUN sudo apt-get install -y gcc python-dev libgit2-dev

COPY ./ /opt/evolution-master
