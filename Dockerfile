FROM debian:stretch

COPY ./ /opt/evolution-master
WORKDIR /opt/evolution-master
RUN apt-get update
RUN apt-get install -y gcc python-dev python-pip libgit2-dev python-cffi python-pygit2
RUN pip install -r requirements.txt
