//
//  singleNumber.cpp
//  leetcode
//
//  Created by Kael on 15/8/21.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

class Solution {
public:
    int singleNumber(int A[], int n) {
        int one=0, two=0, three=0;
        for(int i=0; i<n; i++){
            two |= one&A[i];
            one^=A[i];
            three=one&two;
            one&= ~three;
            two&= ~three;
        }
        return one;
    }
};