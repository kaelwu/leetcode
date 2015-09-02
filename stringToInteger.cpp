//
//  stringToInteger.cpp
//  leetcode
//
//  Created by Kael on 15/5/8.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <string>
#include <math.h>


class Solution {
public:
    int myAtoi(std::string str) {
        long len = str.length();
        long res = 0;
        bool flag = true;
        bool error = false;
        long i;
        int c = 0;
        for (i = 0; i < len; i ++) {
            if (str[i] == '-') {
                flag = false;
                c += 1;
                continue;
            } else if (str[i] == '+') {
                flag += true;
                c += 1;
                continue;
            } else if (str[i] > '9' || str[i] < '0') {
                if (res) {
                    error = true;
                    break;
                } else {
                    continue;
                }
            }
            int num = str[i] - '0';
            res += pow(10, len - i -1) * num;
        }
        res = res / pow(10, len - i);
        if (!flag) {
            return 0 - (int)res;
        }
        if (error) {
            return 0;
        }
        if (res > 2147483647) {
            return 2147483647;
        } else if (res < -2147483648) {
            return -2147483648;
        } else {
            return (int)res;
        }
    }
};