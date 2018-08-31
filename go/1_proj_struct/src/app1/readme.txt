

import and package rule:


import "....../pkg_name"

....../pkg_name: is the directory starting from $GOPATH/src

in the ...../pkg_name: directory.   There can be multiple files.   The file name does not matter.   Inside the file, there are should be:
package pkg_name

The exported fucntion and type should be uper case started.

names between different file should not be duplicated.

We can call reference function between files.



This example demostrates:
* from main, use different package (lib_greet, lib_book)
* in same package (lib_greet), use function from different file.
* in one sub package, use another package (lib_book/book1)
