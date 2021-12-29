# [10. Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/)


## 题目

Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

    '.' Matches any single character.​​​​
    '*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

**Example 1:**
```
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

**Example 2:**
```
Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

**Example 3:**
```
Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

**Example 4:**
```
Input: s = "aab", p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
```

**Example 5:**
```
Input: s = "mississippi", p = "mis*is*p*."
Output: false
```
 
**Constraints:**

 - 1 <= s.length <= 20
 - 1 <= p.length <= 30
 - s contains only lowercase English letters.
 - p contains only lowercase English letters, '.', and '*'.
 - It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.


## Solution

### Approach 1: Recursion

If there were no Kleene stars (the * wildcard character for regular expressions), the problem would be easier - we simply check from left to right if each character of the text matches the pattern.

When a star is present, we may need to check many different suffixes of the text and see if they match the rest of the pattern. A recursive solution is a straightforward way to represent this relationship.

### Approach 2: Dynamic Programming

#### Intuition

As the problem has an optimal substructure, it is natural to cache intermediate results. We ask the question dp(i, j)\text{dp(i, j)}dp(i, j): does text[i:]\text{text[i:]}text[i:] and pattern[j:]\text{pattern[j:]}pattern[j:] match? We can describe our answer in terms of answers to questions involving smaller strings.

#### Algorithm

We proceed with the same recursion as in Approach 1, except because calls will only ever be made to match(text[i:], pattern[j:]), we use dp(i, j)\text{dp(i, j)}dp(i, j) to handle those calls instead, saving us expensive string-building operations and allowing us to cache the intermediate results.

