git clone https://github.com/GoogleCloudPlatform/python-docs-samples
cd python-docs-samples/appengine/standard_python37/hello_world

virtualenv --python python3 ~/envs/hello_world
source ~/envs/hello_world/bin/activate
pip install -r requirements.txt


gcloud app create
gcloud app deploy app.yaml  --project yike-cloud

