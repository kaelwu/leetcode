//
//  trailingZeroes.cpp
//  leetcode
//
//  Created by Kael on 15/5/11.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <math.h>

class Solution {
public:
    int trailingZeroes(int n) {
        double result;
        for (int i = 0 ; i < n; i++) {
            if (i/2 == 0 ) {
                result += 0.5;
            } else if (i/5 == 0) {
                result += 0.5;
            } else if (i / 10 == 0) {
                result += 1;
            }
        }
        int tmp = ceil(result);
        return tmp;
    }
};