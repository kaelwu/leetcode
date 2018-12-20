//
//  reverseList.cpp
//  leetcode
//
//  Created by Kael on 15/5/7.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

/**
 * Definition for singly-linked list.
 */
  struct ListNode {
      int val;
      ListNode *next;
      ListNode(int x) : val(x), next(NULL) {}
  };

class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        ListNode* tmp1 = nullptr;
        ListNode* tmp2 = nullptr;
        if (head == NULL || head->next == NULL) {
            return head;
        }
        while(head->next != NULL) {
            if (tmp1 == NULL) {
                tmp1 = new ListNode(head->val);
            }
            tmp2 = new ListNode(head->next->val);
            tmp2->next = tmp1;
            head = head->next;
            tmp1 = tmp2;
        }
        return tmp2;
    }
};