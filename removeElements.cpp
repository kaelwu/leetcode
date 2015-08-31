//
//  removeElements.cpp
//  leetcode
//
//  Created by Kael on 15/8/14.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

int removeElement(int* nums, int numsSize, int val) {
    int total = 0;
    for (int i = 0; i < numsSize; i ++) {
        if (nums[i] - val == 0) {
            total ++;
        }
    }
    return numsSize - total;
}

