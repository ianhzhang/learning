
from distutils.core import setup
from distutils.extension import Extension
from Cython.Distutils import build_ext


ext_modules = [
                Extension("myLib",  ['py_src/myLib.py']),
                Extension("myLib1", ['py_src/myLib1.py'])
              ]

setup(
  name = 'MyLib',
  cmdclass = {'build_ext': build_ext},
  ext_modules = ext_modules
)
