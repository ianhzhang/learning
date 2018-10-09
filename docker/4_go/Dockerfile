FROM ubuntu:16.04
RUN apt-get update
RUN apt-get install -y --no-install-recommends python

RUN mkdir -p /var/www/server/static

COPY fileserver /var/www/server/
COPY ./static/* /var/www/server/static/
EXPOSE 8000
#CMD ["/var/www/server/fileserver"]