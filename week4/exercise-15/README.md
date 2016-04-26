5 Advanced Problem: Points and Segments

(Recall that advanced problems are not covered in the video lectures and require additional ideas to be solved. We therefore strongly recommend you start solving these problems only when you are done with the basic problems.)
Problem Introduction
The goal in this problem is given a set of segments on a line and a set of points on a line, to count, for each point, the number of segments which contain it.
Problem Description
Task. In this problem you are given a set of points on a line and a set of segments on a line. The goal is to compute, for each point, the number of segments that contain this point.
Input Format. The first line contains two non-negative integers s and p defining the number of segments and the number of points on a line, respectively. The next s lines contain two integers ai,bi defining the i-th segment [ai, bi]. The next line contains p integers defining points x1, x2, . . . , xp.
Constraints. 1≤s,p≤50000;−108 ≤ai ≤bi ≤108 forall0≤i<s;−108 ≤xj ≤108 forall0≤j<p. Output Format. Output p non-negative integers k0, k1, . . . , kp−1 where ki is the number of segments which
contain xi. More formally,
ki =|{j:aj ≤xi ≤bj}|. Time Limits. C: 3.0 sec, C++: 3.0 sec, Java: 4.5 sec, Python: 15.0 sec.
Memory Limit. 64Mb.
