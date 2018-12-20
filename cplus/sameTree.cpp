//
//  sameTree.cpp
//  leetcode
//
//  Created by Kael on 15/8/31.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    bool isSameTree(TreeNode p, TreeNode q) {
        return true;
    }
};