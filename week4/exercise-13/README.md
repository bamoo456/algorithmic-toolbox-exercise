3 Problem: Sorting: 3-Way Partition Problem Introduction

The goal in this problem is to redesign a given implementation of the randomized quick sort algorithm so that it works fast even on sequences containing many equal elements.
Problem Description
Task. To force the given implementation of the quick sort algorithm to efficiently process sequences with few unique elements, your goal is replace a 2-way partition with a 3-way partition. That is, your new partition procedure should partition the array into three parts: < x part, = x part, and > x part.
Input Format. The first line of the input contains an integer n. The next line contains a sequence of n integers a0, a1, . . . , an−1.
Constraints. 1≤n≤105;1≤ai ≤109 forall0≤i<n.
Output Format. Output this sequence sorted in non-decreasing order.
Time Limits. C: 2 sec, C++: 2 sec, Java: 3 sec, Python: 10 sec.
Memory Limit. 64Mb.
