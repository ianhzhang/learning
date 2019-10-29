#include<stdio.h>

double test(double* a, int len) {
    printf("Hello World 444\n");
    double sum = 0.0;
    for(int i=0; i<len; i++) {
        printf("%d: %lf\n", i, a[i]);
        sum = sum + a[i];
    }
    return sum;
}
