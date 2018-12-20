//
//  deleteNodeFromList.cpp
//  leetcode
//
//  Created by Kael on 15/8/14.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>


// Definition for singly-linked list.
struct ListNode {
    int val;
    struct ListNode *next;
};
 
void deleteNode(struct ListNode* node) {
    while (node->next) {
        if (node->next->val == 3) {
            node->next = node->next->next;
        }
    }
}