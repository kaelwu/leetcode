//
//  removeNthNodeFromEndofList.cpp
//  leetcode
//
//  Created by Kael on 15/5/8.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

  struct ListNode {
      int val;
      ListNode *next;
      ListNode(int x) : val(x), next(NULL) {}
  };

class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* tmp = nullptr;
        ListNode* prev = nullptr;
        if (head || head->val == n) {
            return tmp;
        }
        tmp = head;
        while (tmp) {
            if (tmp->val == n) {
                prev->next = tmp->next;
            }
            prev = tmp;
            tmp = tmp->next;
        }
        return tmp;
    }
};
