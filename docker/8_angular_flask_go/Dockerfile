FROM ubuntu:16.04
MAINTAINER Ian Zhang
# 1. nginx

# 1.1 Copy
RUN mkdir -p /var/www/html
COPY ./client/my-client/dist/* /var/www/html

# 1.2 Install
RUN apt-get update
RUN apt-get install -y --no-install-recommends nginx 

# 1.3 Config
RUN rm /etc/nginx/sites-enabled/default
COPY ./config/my_nginx.conf /etc/nginx/sites-available/
RUN ln -s /etc/nginx/sites-available/my_nginx.conf /etc/nginx/sites-enabled/my_nginx.conf
RUN echo "daemon off;" >> /etc/nginx/nginx.conf

# 2. Python Flask server

# 2.1 Copy
RUN mkdir -p /var/www/flaskserver/.not_remove
COPY ./flaskserver /var/www/flaskserver

# 2.2 Install
RUN apt-get install -y --no-install-recommends python2.7 python-pip supervisor
RUN pip install --upgrade pip setuptools wheel
RUN pip install -r /var/www/flaskserver/requirements.txt

# 3. Go server
RUN mkdir -p /var/www/goserver
COPY ./goserver/main /var/www/goserver/
#CMD ["/var/www/goserver/main"]


# 3. Start with Supervisord
COPY ./config/flask_supervisord.conf /etc/supervisor/conf.d/flask.conf
COPY ./config/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

CMD ["/usr/bin/supervisord"]

