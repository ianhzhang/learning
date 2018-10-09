FROM nginx
# install ping for debug
RUN apt-get update
RUN apt-get install -y --no-install-recommends iputils-ping

COPY my-client/dist/my-client/ /usr/share/nginx/html

