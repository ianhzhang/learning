
cdef extern from "c_src/test.h":
    void test()


cpdef pytest():
    print("in pybridge.pyx")
    test()