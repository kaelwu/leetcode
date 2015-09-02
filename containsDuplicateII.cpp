//
//  containsDuplicateII.cpp
//  leetcode
//
//  Created by Kael on 15/9/1.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <vector>
#include <map>

class Solution {
public:
    bool containsNearbyDuplicate(std::vector<int>& nums, int k) {
        std::map<int, int> map1;
        std::vector<int>::iterator it;
        std::pair<std::map<int,int>::iterator,bool> ret;
        bool res = false;
        int i = 1;
        for (it = nums.begin(); it != nums.end(); it ++) {
            ret = map1.insert(std::pair<int, int>(*it, i));
            if (ret.second == false) {
                if (i - ret.first->second <= k) {
                    res = true;
                } else {
                    map1.erase(ret.first);
                    map1.insert(std::pair<int, int>(*it, i));
                }
            }
            i ++;
        }
        return res;
    }
};
