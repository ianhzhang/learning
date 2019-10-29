import pybridge as c_api

a = [1.1, 1.2]
r = c_api.pytest(a)

print("In python result is:" + str(r))
