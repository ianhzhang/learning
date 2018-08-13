IZHANG-C02WK04F:b0_struct izhang$ pwd
/Users/izhang/learningProj/go/b0_struct

-----------------------------------------------------------------------------
IZHANG-C02WK04F:b0_struct izhang$ echo $GOPATH
/Users/izhang/learningProj/go/b0_struct


=============================================================================
app1
=============================================================================

-----------------------------------------------------------------------------
Original tree
IZHANG-C02WK04F:b0_struct izhang$ tree
.
├── app1
├── readme.txt
└── src
    └── app1
        ├── lib_greet
        │   ├── greet1.go
        │   └── greet2.go
        └── main.go

-----------------------------------------------------------------------------
import root:
$GOPATH/src

in main.go, we want to import greet1
import "app1/lib_greet"

The function is identified by package (lib_greet).  Not by the file name (greet1 or greet2)
Under same package, greet1 and greet2 should not have same function name.


-----------------------------------------------------------------------------
Bulid and run:
1. cd app1:
go run main.go

2. cd app1:
go build main.go
This will generate main in the same directory.

3. cd $GOPATH
go build app1

This will generate app1 in the same directory as $GOPATH

4. cd $GOPATH
go install app1

IZHANG-C02WK04F:b0_struct izhang$ tree
.
├── bin
│   └── app1
├── readme.txt
└── src
    └── app1
        ├── lib_greet
        │   ├── greet1.go
        │   └── greet2.go
        └── main.go


=============================================================================
app2
=============================================================================
IZHANG-C02WK04F:b0_struct izhang$ tree
.
├── bin
│   ├── app1
│   └── app2
├── pkg
│   └── darwin_amd64
│       └── app1
│           └── lib_greet.a
├── readme.txt
└── src
    ├── app1
    │   ├── lib_greet
    │   │   ├── greet1.go
    │   │   └── greet2.go
    │   └── main.go
    └── app2
        └── main.go

------------------------------------------------------------------------------
go get app2/...

.
├── bin
│   ├── app1
│   └── app2
├── pkg
│   └── darwin_amd64
│       └── app1
│           └── lib_greet.a
├── readme.txt
└── src
    ├── app1
    │   ├── lib_greet
    │   │   ├── greet1.go
    │   │   └── greet2.go
    │   └── main.go
    ├── app2
    │   └── main.go
    └── github.com
        └── gorilla
            └── mux
                ├── AUTHORS
                ├── ISSUE_TEMPLATE.md
                ├── LICENSE
                ├── README.md
                ├── bench_test.go
                ├── context_gorilla.go
                ├── context_gorilla_test.go
                ├── context_native.go
                ├── context_native_test.go
                ├── doc.go
                ├── example_authentication_middleware_test.go
                ├── example_route_test.go
                ├── go.mod
                ├── middleware.go
                ├── middleware_test.go
                ├── mux.go
                ├── mux_test.go
                ├── old_test.go
                ├── regexp.go
                ├── route.go
                └── test_helpers.go

------------------------------------------------------------
IZHANG-C02WK04F:b0_struct izhang$ go install app2
IZHANG-C02WK04F:b0_struct izhang$ tree
.
├── bin
│   ├── app1
│   └── app2
├── pkg
│   └── darwin_amd64
│       └── app1
│           └── lib_greet.a
├── readme.txt
└── src
    ├── app1
    │   ├── lib_greet
    │   │   ├── greet1.go
    │   │   └── greet2.go
    │   └── main.go
    ├── app2
    │   └── main.go
    └── github.com
        └── gorilla
            └── mux
                ├── AUTHORS
                ├── ISSUE_TEMPLATE.md
                ├── LICENSE
                ├── README.md
                ├── bench_test.go
                ├── context_gorilla.go
                ├── context_gorilla_test.go
                ├── context_native.go
                ├── context_native_test.go
                ├── doc.go
                ├── example_authentication_middleware_test.go
                ├── example_route_test.go
                ├── go.mod
                ├── middleware.go
                ├── middleware_test.go
                ├── mux.go
                ├── mux_test.go
                ├── old_test.go
                ├── regexp.go
                ├── route.go
                └── test_helpers.go