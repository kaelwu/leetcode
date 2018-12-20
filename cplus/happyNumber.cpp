//
//  happyNumber.cpp
//  leetcode
//
//  Created by Kael on 15/9/2.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <math.h>
#include <unordered_set>

class Solution {
public:
    bool isHappy(int n) {
        if (n < 1)
            return false;
        if (n == 1)
            return true;
        std::unordered_set<int> showedNums;
        showedNums.insert(n);
        
        while(true)
        {
            int s = 0;
            while(n)
            {
                s += (n % 10) * (n % 10);
                n = n / 10;
            }
            
            if (s == 1)
                return true;
            else if (showedNums.find(s) != showedNums.end())
                return false;
            n = s;
            showedNums.insert(s);
        }
    }
};
