//
//  isomorphicStrings.cpp
//  leetcode
//
//  Created by Kael on 15/5/8.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <string>
#include <vector>

using std::string;
using std::vector;
using std::pair;

class Solution {
public:
    bool isIsomorphic(string s, string t) {
        vector<int> vector1 = {0};
        vector<int> vector2 = {0};
        string tmp = "";
        if (s.length() != t.length()) {
            return false;
        } else {
            for (int i = 0; i < s.length(); i++) {
                if (vector1[s[i]] == 0) {
                    vector1[s[i]] = t[i];
                }
                if (vector2[t[i]] == 0) {
                    vector2[t[i]] = s[i];
                }
                if (vector2[t[i]] != s[i]) {
                    return false;
                }
                tmp.push_back(vector2[t[i]]);
            }
            if (tmp == s) {
                return true;
            }
            return false;
        }
    }
};
