//
//  validParentheses.cpp
//  leetcode
//
//  Created by Kael on 15/9/1.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <string>
#include <stack>

class Solution {
public:
    bool isValid(std::string s) {
        std::stack<char> stack1;
        bool res = true;
        char temp;
        for (int i = 0; i < s.length(); i ++) {
            switch (s[i]) {
                case '(':
                    ;
                case '{':
                    ;
                case '[':
                    stack1.push(s[i]);
                    break;
                case ')':
                    if (stack1.empty()) {
                        res = false;
                        break;
                    }
                    temp = stack1.top();
                    if (temp != '(') {
                        res = false;
                    }
                    break;
                case '}':
                    if (stack1.empty()) {
                        res = false;
                        break;
                    }
                    temp = stack1.top();
                    if (temp != '{') {
                        res = false;
                    }
                    break;
                case ']':
                    if (stack1.empty()) {
                        res = false;
                        break;
                    }
                    temp = stack1.top();
                    if (temp != '[') {
                        res = false;
                    }
                    break;
                default:
                    break;
            }
            if (!res) {
                break;
            }
        }
        if (!stack1.empty()) {
            res = false;
        }
        return res;
    }
};