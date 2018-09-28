FROM ubuntu:16.04
MAINTAINER Ian Zhang

RUN apt-get update

RUN apt-get install -y --no-install-recommends nginx python2.7 python-pip supervisor
RUN rm -rf /var/lib/apt/lists/*

RUN mkdir -p /var/www/html
RUN mkdir -p /var/www/flaskserver/.not_remove

COPY ./client/my-client/dist/* /var/www/html
COPY ./server /var/www/flaskserver


RUN pip install --upgrade pip setuptools wheel
RUN pip install -r /var/www/flaskserver/requirements.txt
##
RUN rm /etc/nginx/sites-enabled/default
COPY ./config/flask_nginx.conf /etc/nginx/sites-available/
RUN ln -s /etc/nginx/sites-available/flask_nginx.conf /etc/nginx/sites-enabled/flask_nginx.conf


RUN echo "daemon off;" >> /etc/nginx/nginx.conf

COPY ./config/flask.conf /etc/supervisor/conf.d/flask.conf
COPY ./config/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

CMD ["/usr/bin/supervisord"]

