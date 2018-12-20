//
//  removeDuplicatesSortedArray.cpp
//  leetcode
//
//  Created by Kael on 15/8/14.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <vector>
#include <map>

int removeDuplicates(int* nums, int numsSize) {
    int index = 1;
    
    if(!numsSize) {
        return 0;
    }
    for(int i = 1; i < numsSize; i++) {
        if(nums[i] != nums[i - 1]) {
            nums[index] = nums[i];
            index++;
        }
    }
    return index;
}