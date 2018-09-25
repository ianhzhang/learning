docker build -t my-ubuntu .
docker save  my-ubuntu -o my-ubuntu.docker

docker run --name my-ubuntu -d -p 80:80 nginx

IZHANG-C02WK04F:2_nginx izhang$ docker container ls
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                NAMES
b6b784d0e385        nginx               "nginx -g 'daemon of…"   8 seconds ago       Up 7 seconds        0.0.0.0:80->80/tcp   my-ubuntu