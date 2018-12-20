//
//  romaNum.cpp
//  leetcode
//
//  Created by Kael on 15/5/8.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <string>
#include <map>

class Solution {
public:
    int romanToInt(std::string s) {
        int result=0;
        
        std::map<char,int> roman;
        roman['I']=1;
        roman['V']=5;
        roman['X']=10;
        roman['L']=50;
        roman['C']=100;
        roman['D']=500;
        roman['M']=1000;
        
        for(int i=s.length()-1;i>=0;i--)
        {
            if(i==s.length()-1)
            {
                result=roman[s[i]];
                continue;
            }
            if(roman[s[i]] >= roman[s[i+1]])
                result+=roman[s[i]];
            else
                result-=roman[s[i]];
        }
        return result;
    }
};