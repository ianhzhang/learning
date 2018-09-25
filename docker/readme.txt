sudo apt install docker.io
sudo docker run hello-world


1. List all images
docker images

docker container ls -a
docker stop <containerid>
docker rm <containerid>

docker images
docker rmi <imageid>



docker rm $(docker ps -a -q)


2. Build and run:
docker build -t my-ubuntu .
docker save  my-ubuntu -o my-ubuntu.docker
docker load -i my-ubuntu.docker
docker run --name my-ubuntu -d -t --restart=always my-ubuntu


3. Command to run
docker container exec -it bd13205cbe9f pwd
docker container exec -it bd13205cbe9f ls /var
docker container exec -it bd13205cbe9f ls /etc
docker container logs bd13205cbe9f


docker stop $(docker ps -a -q)

docker container rm -f $(docker ps -a -q)
docker rmi $(docker images -a -q)
