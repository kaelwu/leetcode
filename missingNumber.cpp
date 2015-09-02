//
//  missingNumber.cpp
//  leetcode
//
//  Created by Kael on 15/9/2.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

int missingNumber(int* nums, int numsSize) {
    int sum = numsSize*(numsSize + 1) / 2, i;
    for (i = 0; i < numsSize; ++i)
        sum -= nums[i];
    return sum;
}