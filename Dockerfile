FROM debian:stretch

WORKDIR /opt/evolution-master
RUN apt-get update
RUN apt-get install -y gcc python-dev python-pip python-virtualenv libgit2-dev
RUN virtualenv venv
RUN source ./venv/bin/activate
RUN pip install cffi
RUN pip install -r requirements.txt

COPY ./ /opt/evolution-master
