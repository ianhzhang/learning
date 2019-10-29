cimport cython
from libc.stdlib cimport malloc, free

cdef extern from "c_src/test.h":
    double test(double* a, int len)


cpdef pytest(a):
    print("in pybridge.pyx")

    cdef double *my_double
    my_double = <double *>malloc(len(a)*cython.sizeof(double))
    for i in xrange(len(a)):
        my_double[i] = a[i]

    cdef int a_len = len(a)
    cdef double r
    r = test(my_double, a_len)

    free(my_double)
    return r