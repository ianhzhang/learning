import numpy as np
import matplotlib.pyplot as plt

print plt.rcParams

fig_size = plt.rcParams["figure.figsize"]
 
# Prints: [8.0, 6.0]
print "Current size:", fig_size
fig_size[0] = 6
fig_size[1] = 9
plt.rcParams["figure.figsize"] = fig_size
print "-----------"

# Data
clust_data = np.random.random((6,3))  # 4x3 
clust_data1 =[[1,2,3],[4,5,6],[7,9,9]]

collabel=("col 1", "col 2", "col 3")


# laylout
fig, axs =plt.subplots(2,1, gridspec_kw = {'height_ratios':[1, 4]}) # top splot 1/3, bottom 2/3



# title for the whole picture
fig.suptitle("Title for whole figure 1 ", fontsize=16, x=0.25, y=0.95)

# subplots gap
plt.subplots_adjust(hspace=0, top=0.85, bottom=0.15)  # start bottom, total is 1.0.  Top is at 0.85, bottom is at 0.15 

plt.text(0.5, 0.5,'Hello',
     horizontalalignment='center',
     transform = axs[0].transAxes)


# top subplot axs[0]
axs[0].plot(clust_data[:,0],clust_data[:,1])


# bottom subplot axs[1]
axs[1].axis('tight')
axs[1].axis('off')

# https://stackoverflow.com/questions/9932072/matplotlib-table-formatting
#https://stackoverflow.com/questions/27972524/matplotlib-row-heights-table-property

the_table = axs[1].table( cellText=clust_data,      # table data  4x3
                          colLabels=collabel,       # label
                          cellLoc='center',
                          colWidths=[0.5,0.25,0.25],    # columan width
                          bbox=[0,0,1,0.5],     # left, bottom, width, height
                          loc='center')
the_table.set_fontsize(18)
the_table.scale(1,4)

## set border of table
for key, cell in the_table.get_celld().items():
    #print cell
    #cell.set_height(0.2)       ## override by bbox
    #cell.set_linewidth(0)
    pass

plt.show()