cd c_api/c_src
gcc -c test.c

ar rcs librcmodule.a test.o
mv librcmodule.a ../

cd ../
./cython_setup3.sh

cd ..
python3 main.py