//
//  palindromeNumber.cpp
//  leetcode
//
//  Created by Kael on 15/9/1.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <stdlib.h>
#include <string>

class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0)
            return false;
        if (x == 0)
            return true;
        
        int base = 1;
        while(x / base >= 10)
            base *= 10;
        
        while(x)
        {
            int leftDigit = x / base;
            int rightDigit = x % 10;
            if (leftDigit != rightDigit)
                return false;
            
            x -= base * leftDigit;
            base /= 100;
            x /= 10;
        }
        
        return true;
    }
};