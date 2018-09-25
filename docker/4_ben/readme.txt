docker build -t my-ubuntu .
docker save  my-ubuntu -o my-ubuntu.docker

docker run --name my-ubuntu -d -t -p 8080:80 my-ubuntu




 docker container rm -f $(docker ps -a -q)
 docker rmi -f $(docker images -a -q)


docker container rm -f $(docker ps -a -q)
docker rmi -f $(docker images -a -q)
