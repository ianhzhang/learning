cimport cython
from libc.stdlib cimport malloc, free

cdef extern from "c_src/test.h":
    ctypedef struct RCResult:
        double R1
        double C1
        double R2
        double C2

    int test(double* a, int len, RCResult* rslt);





cpdef pytest(a, rslt):
    print("in pybridge.pyx")

    cdef double *my_double
    my_double = <double *>malloc(len(a)*cython.sizeof(double))
    for i in xrange(len(a)):
        my_double[i] = a[i]

    cdef int a_len = len(a)
    cdef int r
    r = test(my_double, a_len, rslt)

    free(my_double)
    return r