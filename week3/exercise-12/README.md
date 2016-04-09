5 Problem: Pairwise Distinct Summands Problem Introduction

This is an example of a problem where a subproblem of the corresponding
greedy algorithm is slightly distinct from the initial problem.
Problem Description
Task. The goal of this problem is to represent a given positive integer
n as a sum of as many pairwise distinct positive integers as possible.
That is, to find the maximum k such that n can be written as
a1+a2+···+ak wherea1,...,ak arepositiveintegersandai ̸=aj forall1≤i<j≤k.
Input Format. The input consists of a single integer n.
Constraints. 1 ≤ n ≤ 109.
Output Format. In the first line, output the maximum number k such that
n can be represented as a sum of k pairwise distinct positive integers.
In the second line, output k pairwise distinct positive integers that
sum up to n (if there are many such representations, output any of
them).
Time Limits. C: 1 sec, C++: 1 sec, Java: 1.5 sec, Python: 5 sec.
Memory Limit. 64Mb.

