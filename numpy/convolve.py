
import numpy as np 

x = np.array([1,2,3])
h = np.array ([1, 5])

y={}

for i in range(0, len(x) + len(h) -1):
    y[i] = 0
    for j in range(len(x)):
        #print("i:", i, "j:",j)
        if i-j >=0 and i-j <=len(h) -1:
            #print("=j:", j, "i:",i)
            y[i] = y[i] + x[j]*h[i-j] 

    print("y[",i,"]=",y[i])


print(np.convolve(x,h))