from selenium import webdriver

driver = webdriver.Chrome()

driver.get('http://python.org')


elem_doc = driver.find_element_by_css_selector('#documentation >a')
elem_doc.click()

elem_py3doc = driver.find_element_by_link_text('Python 3.x Docs')
elem_py3doc.click()
