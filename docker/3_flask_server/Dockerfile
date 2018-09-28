FROM ubuntu:16.04
MAINTAINER Ian Zhang

RUN apt-get update

RUN apt-get install -y --no-install-recommends nginx python2.7 python-pip supervisor
RUN rm -rf /var/lib/apt/lists/*

COPY ./server /

RUN pip install --upgrade pip setuptools wheel
RUN pip install -r /requirements.txt



EXPOSE 5000

#CMD ["/usr/bin/supervisord"]

EXPOSE 6080