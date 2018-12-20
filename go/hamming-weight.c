/*************************************************************************
  > File Name: hamming-wight.c
  > Author: ma6174
  > Mail: ma6174@163.com 
  > Created Time: 四  9/ 6 18:15:34 2018
 ************************************************************************/

#include<stdio.h>

int HammingWeight(int n);

int count;

int main() {
    printf("%d", HammingWeight(11));
}

int HammingWeight(int n) {
    int res = 0;
    for(;n!=0;n = n & (n-1)) {
        printf("test%d\n", n);
        res++;
    }
    return res;
}
