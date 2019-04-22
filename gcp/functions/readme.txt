
1. Python:

1.1 Deploy:
gcloud functions deploy hello_get --runtime python37 --trigger-http

1.2 Access:
https://us-central1-saas-rdu.cloudfunctions.net/hello_get

1.3 Delete
gcloud functions delete hello_get 


2. Go:
2.1 Deploy
gcloud functions deploy HelloGet --runtime go111 --trigger-http

2.2 Delete
gcloud functions delete HelloGet 


Question:
gcloud functions deploy:   which file name it searches?  Multiple file?  Package?
https://cloud.google.com/sdk/gcloud/reference/functions/deploy

The current deployment tool takes a simple approach. 
It zips the directory and uploads it. 
This means you (currently) should move or delete node_modules before executing the command 
if you don't wish for them to be included in the deployment package. 
Note that, like lambda, GCF will resolve dependencies automatically.
