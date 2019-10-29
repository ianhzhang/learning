import requests
import os


file_fullpath = "./index.html"
file_basename = os.path.basename(file_fullpath)

multipart_form_data = {
    'myFile': (file_basename, open(file_fullpath, 'rb'))
}

resp = requests.post('http://localhost:8080/upload', files= multipart_form_data)
print(resp.status_code)
print(resp.content)
