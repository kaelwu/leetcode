//
//  26.cpp
//  leetcode
//
//  Created by Kael on 15/8/14.
//  Copyright (c) 2015年 Kael. All rights reserved.
//
#include <stdio.h>
#include <math.h>
#include <string.h>


#include <stdio.h>

int titleToNumber(char* s) {
    int length = strlen(s);
    int result = 0;
    for (int i = 0; i < length; i ++) {
        int tmp = s[length - i - 1] - (int)'A';
        if (i > 0) {
            result += pow(26, i) * (tmp + 1);
        } else {
            result += pow(26, i) + tmp;
        }
    }
    return result;
}