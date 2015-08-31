//
//  powerOfTwo.cpp
//  leetcode
//
//  Created by Kael on 15/8/31.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

bool isPowerOfTwo(int n) {
    int temp = n;
    if (n == 0) {
        return false;
    }
    if (n == 1) {
        return true;
    }
    bool res = false;
    while (temp % 2 == 0) {
        temp = temp / 2;
        if (temp == 1) {
            res = true;
        }
    }
    return res;
}