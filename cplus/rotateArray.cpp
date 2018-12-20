//
//  rotateArray.cpp
//  leetcode
//
//  Created by Kael on 15/5/8.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <vector>

using std::vector;

class Solution {
public:
    void rotate(vector<int>& nums, int k) {
        int tmp = k;
        vector<int> result;
        if (nums.size() < k) {
            for (int i =0 ; i < nums.size(); i++) {
                result.push_back(nums[nums.size() - 1]);
            }
        } else {
            for (int i = 0; i < nums.size(); i++) {
                if (tmp > 0) {
                    result.push_back(nums[nums.size() - tmp]);
                    tmp--;
                } else {
                    result.push_back(nums[i - k]);
                }
            }
        }
        nums = result;
    }
};