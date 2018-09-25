FROM ubuntu:16.04
RUN apt-get update

RUN apt-get install -y --no-install-recommends nginx

RUN mkdir -p /var/www/html

COPY ./client/my-client/dist/* /var/www/html
COPY ./run.sh /


# application server
RUN mkdir -p /var/www/flaskserver/.not_remove
RUN apt-get install -y --no-install-recommends python2.7 python-pip
RUN apt-get install -y --no-install-recommends build-essential python
RUN apt-get install -y --no-install-recommends python-dev
RUN pip install --upgrade pip setuptools
COPY ./requirements.txt /
RUN pip install -r /requirements.txt
COPY ./config/flask.service /etc/systemd/system/
COPY ./config/flask_nginx.conf /etc/nginx/sites-available/
RUN ls -s  /etc/nginx/sites-available/flask_nginx.conf /etc/nginx/sites-enabled/
COPY ./server /var/www/flaskserver
COPY ./config/proj.ini /var/www/flaskserver/
CMD ["/run.sh"]

