FROM ubuntu:16.04
MAINTAINER Ian Zhang

RUN apt-get update

RUN apt-get install -y --no-install-recommends python2.7 python-pip
RUN rm -rf /var/lib/apt/lists/*

COPY ./server /

RUN pip install --upgrade pip setuptools wheel
RUN pip install -r /requirements.txt

# useless EXPOSE
EXPOSE 5000    

CMD ["python", "app.py"]