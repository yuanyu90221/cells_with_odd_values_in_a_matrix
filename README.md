# cells_with_odd_values_in_a_matrix

## 題目解讀：

### 題目來源:
[cells-with-odd-values-in-a-matrix](https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/)

### 原文:
Given n and m which are the dimensions of a matrix initialized by zeros and given an array indices where indices[i] = [ri, ci]. For each pair of [ri, ci] you have to increment all cells in row ri and column ci by 1.

Return the number of cells with odd values in the matrix after applying the increment to all indices.

### 解讀：

給定一個 n by m的矩陣 matrix 初始化所有元素為0 給定個一個 座標陣列 indice

每個 indice[i] = [ri, ci]

每次需要把 matrix 中的 row ri 內每個元素 加1 以及 column ci 內的每個元素 加1

求最後運算後 matrix中包含多少個奇數


## 初步解法:
### 初步觀察:

對於每個座標的 ri, ci 個別代表一條線
            
先個別用 map 去計算有幾個不同x跟y

然後把不同個數的值相乘就可以得到 總共有幾個交點

舉例來說：

Input: n = 2, m = 3, indices = [[0,1],[1,1]]

得到有兩個相同的columns 也就是 x=1 所以 x的個數是1

而有兩個不同 rows 所以y的個數是2


所以總共有 x個數 * y的個數 個交點

其中焦點的值 是  x的值 + 1 = 2 + 1 是odd

除去交點的其他點
每個 rows 代表 m個值 除去 x個數 代表有 m-x 個 x 總點數 = (3 -1) * 2 = 4 因為每個點 都只有 1 加一次 所以是odd

給個 columns 代表 n個值 除去 y個數 代表有 n-y個 x 總點數 = (2 -2) * 2= 0


### 初步設計:

我自己的原本設計:

Given two integers n > 0,  m > 0, and a point array indices

Step 1 Set a map X to store all unique column values

Step 2 Set a map Y to store all unique row  values

Step 3 let xCount = 0, yCount = 0, result = 0, crossCount = 0

Step 4 loop X if xi is odd set xCount = xCount + 1

Step 5 loop Y if yi is odd set yCount = yCount + 1

step 6 result += yCount * (m - len(X))

step 7 result += xCount * (n - len(Y))

step 8 result += (yCount * (len(Y) - xCount) + xCount * (len(X) - yCount))

step 9 return result
## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package odd_cell

func oddCells(n int, m int, indices [][]int) int {
	result := 0
	Xmap := make(map[int]int)
	Ymap := make(map[int]int)
	xCount := 0
	yCount := 0
	for _, point := range indices {
		Xmap[point[1]] += 1
		Ymap[point[0]] += 1
	}
	for _, value := range Xmap {
		if value%2 == 1 {
			xCount += 1
		}
	}
	for _, value := range Ymap {
		if value%2 == 1 {
			yCount += 1
		}
	}
	result += yCount * (m - len(Xmap))
	result += xCount * (n - len(Ymap))
	result += xCount*(len(Ymap)-yCount) + yCount*(len(Xmap)-xCount)
	return result
}

```
## 測資的撰寫
```golang
package odd_cell

import "testing"

func Test_oddCells(t *testing.T) {
	type args struct {
		n       int
		m       int
		indices [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				n:       2,
				m:       3,
				indices: [][]int{{0, 1}, {1, 1}},
			},
			want: 6,
		},
		{
			name: "Example2",
			args: args{
				n:       2,
				m:       2,
				indices: [][]int{{1, 1}, {0, 0}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddCells(tt.args.n, tt.args.m, tt.args.indices); got != tt.want {
				t.Errorf("oddCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record

[golang leetcode 30 day 22thday](https://hackmd.io/mvDZf4DCT7-3GUu_vyWcNg?view)

## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)