FROM debian:stretch

WORKDIR /opt/evolution-master
RUN apt-get update
RUN apt-get install -y gcc python-dev python-pip libgit2-dev
RUN pip install cffi
RUN pip install -r requirements.txt

COPY ./ /opt/evolution-master
