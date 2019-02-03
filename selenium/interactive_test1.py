from selenium import webdriver

driver = webdriver.Chrome()

driver.get('http://stackoverflow.com')

element = driver.find_element_by_name('q')
element.send_keys('selenium webdriver')

from selenium.webdriver.common.keys import Keys
element.send_keys(Keys.ENTER)