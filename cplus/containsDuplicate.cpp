//
//  containsDuplicate.cpp
//  leetcode
//
//  Created by Kael on 15/8/31.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <vector>
#include <map>

class Solution {
public:
    bool containsDuplicate(std::vector<int>& nums) {
        std::vector<int>::iterator it;
        if (nums.empty()) {
            return false;
        }
        std::map<int, int> map1;
        std::pair<std::map<int,int>::iterator,bool> ret;
        for (it = nums.begin(); it != nums.end(); it ++) {
            int num = *it;
            ret = map1.insert(std::pair<int, int>(num, 1));
            if (ret.second == false) {
                return true;
            }
        }
        return false;
    }
};