# [29. Divide Two Integers](https://leetcode.com/problems/divide-two-integers/)


## 题目

Given two integers `dividend` and `divisor`, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing `dividend` by `divisor`.

The integer division should truncate toward zero.

**Example 1:**

    Input: dividend = 10, divisor = 3
    Output: 3

**Example 2:**

    Input: dividend = 7, divisor = -3
    Output: -2

**Note:**

- Both dividend and divisor will be 32-bit signed integers.
- The divisor will never be 0.
- Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31, 2^31 − 1]. For the purpose of this problem, assume that your function returns 2^31 − 1 when the division result overflows.


## 题目大意

给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。返回被除数 dividend 除以除数 divisor 得到的商。

说明:

- 被除数和除数均为 32 位有符号整数。
- 除数不为 0。
- 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。


## 解题思路

### 人类计算除法
当我们在计算51/3=17，抛开9*9乘法表。

1. 从被除数的最高位5开始，从0-9选一个数，使得5-i*3>=0且使5-（i+1)*3<0。我们选择了1. 余数为2.
2. 将余数*10+1=21,继续从0-9中选一个数，使得21-3*i>=0且使5-（i+1)*3<0，我们选择了7.
3. 由此，我们找到了答案17。

### 计算机计算除法
计算机计算除法的过程与人类计算的过程很类似，只是选择范围变成了0或1.
还以51/3为例说明（51：110011；3:11）

1. 从第一位开始为1，小于11，结果位置0；余数为1
2. 从第二位开始，余数*2+1=11，等于11，结果位置1，余数为0；
3. 从第三、四位开始，余数*2+0=0<011,结果位置0，余数为0
4. 从第5位开始，余数*2+1=1<11，结果置0，余数为1
5. 从第6位开始，余数*2+1=11=11，结果置1，余数为0.

see: https://blog.csdn.net/zdavb/article/details/47108505
