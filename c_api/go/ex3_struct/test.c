#include<stdio.h>
#include "test.h"


RCResult test(double* a, int len) {
    printf("Hello World 444\n");
    double sum = 0.0;
    for(int i=0; i<len; i++) {
        printf("%d: %lf\n", i, a[i]);
        sum = sum + a[i];
    }
    RCResult rslt;
    rslt.R1 = 1.1;
    rslt.C1 = 1.2;
    rslt.R2 = 1.3;
    rslt.C2 = 1.4;
    return rslt;
}
