# Description

The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```json
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

**Example 1:**

>Input: s = "PAYPALISHIRING", numRows = 3\
>Output: "PAHNAPLSIIGYIR"

**Example 2:**

>Input: s = "PAYPALISHIRING", numRows = 4\
>Output: "PINALSIGYAHRPI"

Explanation:

```json
P     I     N   [0 = i + n * 2 - 2 - 2 * 0, 6 = i + n * 2 - 2]
A   L S   I G   [4 = i + n * 2 - 2 - 2 * 1, 6 = i + n * 2 - 2]
Y A   H R       [2 = i + n * 2 - 2 - 2 * 2, 6 = i + n * 2 - 2]
P     I         [0 = i + n * 2 - 2 - 2 * 3, 6 = i + n * 2 - 2]
```

Example 3:

>Input: s = "A", numRows = 1\
>Output: "A"

**Constraints:**

`1 <= s.length <= 1000`
`s` consists of English letters (lower-case and upper-case), `','` and `'.'`.

`1 <= numRows <= 1000`
