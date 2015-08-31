//
//  uglyNum.cpp
//  leetcode
//
//  Created by Kael on 15/8/19.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

int test(int num) {
    int temp = num;
    if ((double)num / 2 - (int)(num / 2) == 0) {
        return test(num / 2);
    } else if ((double)num / 3 - (int)(num / 3) == 0) {
        return test(num / 3);
    } else if ((double)num / 5 - (int)(num / 5) == 0) {
        return test(num /5);
    }
    return temp;
}

bool isUgly(int num) {
    int temp = test(num);
    if (temp == 1) {
        return true;
    } else {
        return false;
    }
    
}
