1 Problem: Primitive Calculator Problem Introduction

You are given a primitive calculator that can perform the following three operations with the current num- ber x: multiply x by 2, multiply x by 3, or add 1 to x. Your goal is given a positive integer n, find the minimum number of operations needed to obtain the number n starting from the number 1.
Problem Description
Task. Given an integer n, compute the minimum number of operations needed to obtain the number n starting from the number 1.
Input Format. The input consists of a single integer 1 ≤ n ≤ 106.
Output Format. In the first line, output the minimum number k of operations needed to get n from 1. In the second line output a sequence of intermediate numbers. That is, the second line should contain positiveintegersa0,a2,...,ak−1 suchthata0 =1,ak−1 =nandforall0≤i<k−1,ai+1 isequalto either ai + 1, 2ai, or 3ai. If there are many such sequences, output any one of them.
Time Limits. C: 1 sec, C++: 1 sec, Java: 1.5 sec, Python: 5 sec. 3 sec, Ruby: 3 sec, Scala: 3 sec.
Memory Limit. 128Mb.

