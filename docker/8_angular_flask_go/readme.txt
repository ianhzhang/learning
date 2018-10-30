
dockerclean
dockerbuild
docker run --name my-ubuntu -d -t -p 4200:80 my-ubuntu

external port 4200
internal port 80

cd client/my-client
npm i
ng serve

http://localhost:4200/api/hello
will get 
{"rslt":"hello World 1"}

http://localhost:4200/
will access internal ui.



Summary:
dockerclean
dockerbuild
docker run --name my-ubuntu -d -t -p 4200:80 my-ubuntu
http://localhost:4200/