
from distutils.core import setup
from distutils.extension import Extension
from Cython.Distutils import build_ext
from Cython.Build.Dependencies import cythonize

ext_modules = [
        Extension(name="pybridge", 
                  sources=["./pybridge.pyx"],
                  libraries=["rcmodule"],
                  library_dirs=["."]),
]

setup(
  name = 'pybridge',
  cmdclass = {'build_ext': build_ext},
  ext_modules=cythonize(ext_modules, compiler_directives={'always_allow_keywords': True})
)
