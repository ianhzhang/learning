import numpy as np
import matplotlib.pyplot as plt
plt.figure(0)

axes1 = plt.subplot2grid((3,3), (0,0), colspan=3)
axes2 = plt.subplot2grid((3,3), (1,0), colspan=2)
axes3 = plt.subplot2grid((3,3), (1,2), colspan=1, rowspan=2)
axes4 = plt.subplot2grid((3,3), (2,0), colspan=1)
axes5 = plt.subplot2grid((3,3), (2,1), colspan=1)


axes1.set_facecolor((1.0,0.47,0.42))
axes2.set_facecolor((0.4,0.47,0.12))
axes3.set_facecolor((0.3,0.27,0.42))
axes4.set_facecolor((0.2,0.47,0.32))
axes5.set_facecolor((1.0,0.47,0.82))
plt.suptitle("Demo subplot suptitle")
#plt.title("title")      # in the middle
plt.show()