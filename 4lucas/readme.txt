1. To Build and Run a Docker Container:
cd docker_build
docker build -t netapp-proj .
docker run --name netapp-proj -d -t -p 8085:8085 netapp-proj

Note: Since pytorch needs to download all pre-trained Machine Learning parameters and analyse images, it may take up to 3 mins to bring up the server. 

To run web UI:
http://0.0.0.0:8085


2. To Run Without Docker Container
* cd docker_build/server/
* python3 server_main.py
( You need to have flask, flask_cors, Pillow==5.0.0 pytorch, torchversion installed )

3. To Build UI from Source Code (I have already packaged for you):
* install nodejs/npm
* install angular
* cd src/ui
* ng build --prod
* cp -a dist/ui/* to docker_build/server/static
