//
//  validAnagram.cpp
//  leetcode
//
//  Created by Kael on 15/8/31.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <string>
#include <map>

class Solution {
public:
    bool isAnagram(std::string s, std::string t) {
        if (s.length() != t.length()) {
            return false;
        }
        std::multimap<char, int> map1;
        for (int i = 0; i < s.length(); i ++) {
            map1.insert(std::pair<char, int>(s[i], i));
        }
        for (int j = 0; j < t.length(); j ++) {
            std::multimap<char, int>::iterator iter=map1.find(t[j]);
            if (iter == map1.end()) {
                return false;
            } else {
                map1.erase(iter);
            }
        }
        if (!map1.empty()) {
            return false;
        }
        return true;
    }
};