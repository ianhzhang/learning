rm -rf build
rm *.c
rm *.so

python3 setup.py build_ext --inplace
cp *.so bin/

# clean
rm -rf build
rm py_src/*.c
rm *.so
