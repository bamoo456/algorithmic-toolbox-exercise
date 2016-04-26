5 Advanced Problem: Longest Common Subsequence of Three Sequences

(Recall that advanced problems are not covered in the video lectures and require additional ideas to be solved. We therefore strongly recommend you start solving these problems only when you are done with the basic problems.)
Problem Introduction
In this problem, your goal is to compute the length of a longest common subsequence of three sequences.
Problem Description
Task. Given three sequences A = (a1,a2,...,an), B = (b1,b2,...,bm), and C = (c1,c2,...,cl), find the length of their longest common subsequence, i.e., the largest non-negative integer p such that there existindices1≤i1 <i2 <···<ip ≤n,1≤j1 <j2 <···<jp ≤m,1≤k1 <k2 <···<kp ≤lsuch that ai1 = bj1 = ck1,...,aip = bjp = ckp
Input Format. First line: n. Second line: a1, a2, . . . , an. Third line: m. Fourth line: b1, b2, . . . , bm. Fifth line: l. Sixth line: c1,c2,...,cl.
Constraints. 1 ≤ n,m,l ≤ 100; −109 < ai,bi,ci < 109.
Output Format. Output p.
Time Limits. C: 2 sec, C++: 2 sec, Java: 3 sec, Python: 10 sec. sec, Ruby: 6 sec, Scala: 6 sec.
Memory Limit. 128Mb.
