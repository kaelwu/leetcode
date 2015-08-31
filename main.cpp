//
//  main.cpp
//  leetcode
//
//  Created by Kael on 15/5/7.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <math.h>

#include <stdio.h>

int hammingWeight(int n) {
    int num = 0;
    while (n / 2) {
        if (n % 2 == 1) {
            num ++;
        }
        n = n / 2;
    }
    return num;
}

int main() {
    int a = hammingWeight(11);
    printf("%d", a);
    return 0;
}
