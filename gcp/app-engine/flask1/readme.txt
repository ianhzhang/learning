pip install -t lib -r requirements.txt 

Test:
dev_appserver.py .
dev_appserver.py app.yaml 

http://localhost:8080/
http://localhost:8000



The difference between this one and flas0 is that:

create appengine_config.py
from google.appengine.ext import vendor

# Add any libraries installed in the "lib" folder.
vendor.add('lib')


We don't need:
import sys
import os

sys.path.insert(1, os.path.join(os.path.abspath('.'), 'lib'))

