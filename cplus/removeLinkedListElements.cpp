//
//  removeLinkedListElements.cpp
//  leetcode
//
//  Created by Kael on 15/5/8.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>
struct ListNode {
         int val;
         ListNode* next;
         ListNode(int x) : val(x), next(NULL) {}
     };

class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        ListNode* tmp = head;
        ListNode* newHead = nullptr;
        ListNode* prev = nullptr;
        ListNode* node = nullptr;
        while(tmp) {
            if (tmp->val == val) {
                tmp = tmp->next;
                continue;
            }
            node = new ListNode(tmp->val);
            if (prev) {
                prev->next = node;
            } else {
                newHead = node;
            }
            prev = node;
            tmp = tmp->next;
        }
        return newHead;
    }
};