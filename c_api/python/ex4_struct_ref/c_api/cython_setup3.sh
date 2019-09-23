rm -rf build
rm -rf bin
rm -rf *.c

python3 setup.py build_ext --inplace
mkdir bin
#cp src/main.py bin/
cp *.so ../
mv *.so bin/
cp main.py bin/



rm -rf build
rm -rf *.c
