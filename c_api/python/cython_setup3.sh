rm -rf build
rm -rf bin
rm -rf *.so
rm -rf src/*.pyc
rm -rf src/*.c
rm -rf src/*.so
rm -rf src/algorithm_module/*.so
rm -rf src/algorithm_module/*.c
rm -rf src/algorithm_module/*.pyc

rm -rf src/deviceDriver/*.so
rm -rf src/deviceDriver/*.c
rm -rf src/deviceDriver/*.pyc

mkdir -p bin/algorithm_module
mkdir -p bin/deviceDriver


python3 setup.py build_ext --inplace

cp src/main.py bin/
mv *.so bin/

mv algorithm_module/*.so bin/algorithm_module/
cp src/algorithm_module/__init__.py bin/algorithm_module/

mv deviceDriver/*.so bin/deviceDriver/
cp src/deviceDriver/__init__.py bin/deviceDriver/

cp src/config.ini bin/
cp -a src/output bin/

rm -rf build
rm -rf src/*.c
rm -rf src/*.pyc

rm -rf src/algorithm_module/*.c
rm -rf src/algorithm_module/*.pyc
rm -rf algorithm_module

rm -rf src/deviceDriver/*.c
rm -rf src/deviceDriver/*.pyc
rm -rf deviceDriver