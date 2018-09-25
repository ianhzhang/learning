sudo docker build -t my-ubuntu .
docker save  my-ubuntu -o my-ubuntu.docker

sudo docker run --name my-ubuntu -d -p 8080:80 my-ubuntu




sudo docker container rm -f $(sudo docker ps -a -q)
sudo docker rmi $(sudo docker images -a -q)