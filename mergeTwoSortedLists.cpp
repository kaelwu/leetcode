//
//  mergeTwoSortedLists.cpp
//  leetcode
//
//  Created by Kael on 15/5/8.
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
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode* l3 = nullptr;
        ListNode* prev = nullptr;
        ListNode* tmp = nullptr;
        while (l1 || l2) {
            int val1 = 9999999, val2 = 99999999;
            if (l1) {
                val1 = l1 -> val;
            }
            if (l2) {
                val2 = l2 ->val;
            }
            if (val1 <= val2) {
                tmp = new ListNode(val1);
                l1 = l1->next;
            } else {
                tmp = new ListNode(val2);
                l2 = l2->next;
            }
            if (!prev) {
                l3 = tmp;
            } else {
                prev->next = tmp;
            }
            prev = tmp;
        }
        return l3;
    }
};