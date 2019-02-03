
# Test case description
# 1. Open a browser
# 2. Go to python.org
# 3. Web page title contains word python

from selenium import webdriver

driver = webdriver.Chrome()

driver.get('http://python.org')
assert 'python' in driver.title.lower()
driver.close()

# python