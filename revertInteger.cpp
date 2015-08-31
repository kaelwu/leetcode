//
//  revertInteger.cpp
//  leetcode
//
//  Created by Kael on 15/8/14.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <math.h>

int reverse(int x) {
    int revert = 0;
    int *arr;
    int i = 0;
    while (x / 10 != 0) {
        int tmp = x/10;
        *(arr + i) = tmp;
        i ++;
    }
    for (int j = 0; j <= i; j++) {
        revert += arr[i - j] * pow(10, j);
    }
    if (x < 0) {
        revert = -revert;
    }
    return revert;
};