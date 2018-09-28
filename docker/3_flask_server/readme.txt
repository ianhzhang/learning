docker build -t my-ubuntu .
docker save  my-ubuntu -o my-ubuntu.docker

docker run --name my-ubuntu -d -t -p 8080:80 my-ubuntu




 docker container rm -f $(docker ps -a -q)
 docker rmi -f $(docker images -a -q)


docker container rm -f $(docker ps -a -q)
docker rmi -f $(docker images -a -q)


http://localhost:8080/hello
docker run --name my-ubuntu -d -t -p 8080:5000 my-ubuntu; docker container ls



alias dockerrmi="docker rmi -f \$(docker images -a -q)"
alias dockerrmc="docker container rm -f \$(docker ps -a -q)"
alias dockerbuild="docker build -t my-ubuntu ."
alias dockerrun="docker run --name my-ubuntu -d -t -p 8080:80 my-ubuntu;docker container ls"
alias dockerclean="docker container rm -f \$(docker ps -a -q); docker rmi -f \$(docker images -a -q)"
function dockerbash() {
   echo $1
   echo "xxx"
   docker container exec -it $1 /bin/bash
}