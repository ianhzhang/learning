FROM python:3-onbuild
# install ping for debug
RUN apt-get update
RUN apt-get install -y --no-install-recommends iputils-ping
COPY . /usr/src/app
CMD ["python", "app.py"]