
from distutils.core import setup
from distutils.extension import Extension
from Cython.Distutils import build_ext

sourcefiles = ['py_src/python2cAPI.pyx']
ext_modules = [Extension("python2cAPI", 
                          sourcefiles
                          )]

setup(
  name = 'python2cAPI.py',
  cmdclass = {'build_ext': build_ext},
  ext_modules = ext_modules
)
