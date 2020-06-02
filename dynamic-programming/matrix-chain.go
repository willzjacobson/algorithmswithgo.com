/*
Matrix multiplication is associative. That is, when you have a A1 * A2 * A3, you can calculate the result via (A1 * A1) * A3, or via A1 * (A1 * A3).
The order you pick can have a huge effect on how many scalar operations you have to perform in total.
When you multiply a pxq matrix by a qxr matrix, you perform p*q*r scalar operations.

Imagine you're mutiplying 3 matrices:
	A1: 10 x 100
	A2: 100 x 5
	A3: 5 x 50

Multiplying:
	(A1 * A1) * A3
requires 10*100*5 + 10*5*50 = 7500 scalar operations, while multiplying: 
	A1 * (A1 * A3)
requires 100*5*50 + 10*100*50 = 75000 scalar operations. That's 10x more!

In the matric chain multiplication problem, we're not actually multiplying matrices. Instead, we're figuring out what is the optimal way to paranthesize them for multiplication.
*/