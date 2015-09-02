//
//  compareVersionNumbers.cpp
//  leetcode
//
//  Created by Kael on 15/9/1.
//  Copyright (c) 2015年 Kael. All rights reserved.
//

#include <stdio.h>

int compareVersion(char* version1, char* version2) {
    const char *d = ".";
    char *p, *q;
    p = strtok(version1, d);
    q = strtok(version2, d);
    while (p && q) {
        if (p > q) {
            return 1;
        } else if (p < q) {
            return -1;
        }
        p=strtok(NULL, d);
        q=strtok(NULL, d);
    }
    return 0;
}