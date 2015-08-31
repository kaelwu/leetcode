//
//  removeDuplicatesFromSortedList.cpp
//  leetcode
//
//  Created by Kael on 15/8/31.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

struct ListNode* deleteDuplicates(struct ListNode* head) {
    int temp = head->val;
    struct ListNode* res = head, *tmpNode = head;
    if (!head) {
        return head;
    }
    while (tmpNode->next) {
        if (temp == tmpNode->val) {
            tmpNode = tmpNode->next;
            continue;
        } else {
            temp = tmpNode->val;
            tmpNode = tmpNode->next;
            struct ListNode* next;
            next->val = temp;
            res->next = next;
        }
    }
    return res;
}