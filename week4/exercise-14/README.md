4 Advanced Problem: Number of Inversions

(Recall that advanced problems are not covered in the video lectures and require additional ideas to be solved. We therefore strongly recommend you start solving these problems only when you are done with the basic problems.)
Problem Introduction
An inversion of a sequence a0,a1,...,an−1 is a pair of indices 0 ≤ i < j < n such that ai > aj. The number of inversions of a sequence in some sense measures how close the sequence is to being sorted. For example, a sorted (in non-descending order) sequence contains no inversions at all, while in a sequence sorted in descending order any two elements constitute an inversion (for a total of n(n − 1)/2 inversions).
Problem Description
Task. The goal in this problem is to count the number of inversions of a given sequence.
Input Format. The first line contains an integer n, the next one contains a sequence of integers
a0,a1,...,an−1.
Constraints. 1≤n≤105,1≤ai ≤109 forall0≤i<n.
Output Format. Output the number of inversions in the sequence.
Time Limits. C: 3 sec, C++: 3 sec, Java: 4.5 sec, Python: 15 sec.
Memory Limit. 64Mb.
