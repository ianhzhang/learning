FROM ubuntu:16.04
RUN apt-get update

RUN apt-get install -y --no-install-recommends nginx

RUN mkdir -p /var/www/html

COPY ./static /var/www/html
COPY ./run.sh /
CMD ["/run.sh"]