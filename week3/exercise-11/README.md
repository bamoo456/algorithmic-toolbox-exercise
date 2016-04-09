4 Problem: Covering Segments by Points Problem Introduction

You are given a set of segments on a line and your goal is to mark as
few points on a line as possible so that each segment contains at least
one marked point.
Problem Description
Task. Given a set of n segments {[a0,b0],[a1,b1],...,[an−1,bn−1]} with
integer coordinates on a line, find the minimum number m of points such
that each segment contains at least one point. That is, find a set of
integers X of the minimum size such that for any segment [ai,bi] there
is a point x ∈ X such thatai ≤x≤bi.
Input Format. The first line of the input contains the number n of
segments. Each of the following n lines contains two integers ai and bi
(separated by a space) defining the coordinates of endpoints of the i-th
segment.
Constraints. 1≤n≤100;0≤ai ≤bi ≤109 forall0≤i<n.
Output Format. Output the minimum number m of points on the first line
and the integer coordinates of m points (separated by spaces) on the
second line. You can output the points in any order. If there are many
such sets of points, you can output any set. (It is not difficult to see
that there always exist a set of points of the minimum size such that
all the coordinates of the points are integers.)
Time Limits. C: 1 sec, C++: 1 sec, Java: 1.5 sec, Python: 5 sec.
Memory Limit. 64Mb.

