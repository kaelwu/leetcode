//
//  newPow.cpp
//  leetcode
//
//  Created by Kael on 15/9/2.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

class Solution {
public:
    double myPow(double x, int n) {
        double sum = 1;
        int plus = 1;
        if (n < 0) {
            n = -n;
            plus = 0;
        } else if (n == 0) {
            return 1;
        }
        for (int i = 0; i < n; i ++) {
            sum *= x;
        }
        if (!plus) {
            sum = 1 / sum;
        }
        return sum;
    }
};