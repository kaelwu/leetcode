//
//  lowestCommonAncestorOfBST.cpp
//  leetcode
//
//  Created by Kael on 15/8/19.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>


/**
 * Definition for a binary tree node.
 */
  struct TreeNode {
      int val;
      struct TreeNode *left;
      struct TreeNode *right;
  };
 
struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q) {
    return root;
}