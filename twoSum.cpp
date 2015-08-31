//
//  twoSum.cpp
//  leetcode
//
//  Created by Kael on 15/5/7.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
#include <vector>
#include <map>

using std::vector;
using std::map;
using std::pair;

class Solution {
public:
    vector<int> twoSum(vector<int> &numbers, int target) {
        int i;
        vector<int> results;
        map<int, int> hmap;
        for(i=0; i<numbers.size(); i++){
            if(!hmap.count(numbers[i])){
                hmap.insert(pair<int, int>(numbers[i], i));
            }
            if(hmap.count(target-numbers[i])){
                int n=hmap[target-numbers[i]];
                if(n<i){
                    results.push_back(n+1);
                    results.push_back(i+1);
                    //cout<<n+1<<", "<<i+1<<endl;
                    return results;
                }
                
            }
        }
        return results;
    }
};