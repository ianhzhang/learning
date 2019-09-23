#include<stdio.h>
#include "test.h"


int test(double* a, int len, RCResult* rslt) {
    printf("Hello World 444\n");
    double sum = 0.0;
    for(int i=0; i<len; i++) {
        printf("%d: %lf\n", i, a[i]);
        sum = sum + a[i];
    }
    printf("=x=x=x=x=x=x\n");
    printf("%lf\n", rslt->R2);
    
    rslt->R1 = 1.1;
    rslt->C1 = 1.2;
    rslt->R2 = 1.3;
    rslt->C2 = 1.4;
    return 1;
}
