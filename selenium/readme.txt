1. Python virtual env
sudo apt install python3-pip
sudo apt-get install python3-venv

1.1 Create
python3 -m venv ian_p3
(in current directory, create ian_p3)

1.2 Activate
source ian_p3/bin/activate
(which python
python --version
)

1.3 Deactive
deactivate

2. install python selenium
pip install selenium
(ian_p3/lib/python3.6/site-packages/selenium/)


3. Install ChromeDriver
http://chromedriver.chromium.org/
download latest(2.45) chromedriver_linux64.zip
google-chrome --version
sudo unzip -d /usr/local/bin ~/Downloads/chromedriver_linux64.zip 
ls /usr/local/bin/chromedriver 


