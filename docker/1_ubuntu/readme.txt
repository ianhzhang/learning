docker network create ihz-n
docker network ls

dockerbuild
docker run --name my-ubuntu1 -d -t --network="ihz-n"  my-ubuntu
docker run --name my-ubuntu2 -d -t --network="ihz-n"  my-ubuntu


docker container exec -it my-ubuntu1 /bin/bash
ping my-ubuntu2

dockerclean
docker network rm ihz-n
