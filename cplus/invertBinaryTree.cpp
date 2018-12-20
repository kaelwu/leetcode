//
//  invertBinaryTree.cpp
//  leetcode
//
//  Created by Kael on 15/9/1.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

 struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };

class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        if(root==NULL)
            return NULL;
        TreeNode * ptmpNode = root->left;
        root->left = invertTree(root->right);
        root->right = invertTree(ptmpNode);
        return root;
    }
};