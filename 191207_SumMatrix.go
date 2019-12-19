package main

//给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2)。
//
//
//上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。
//
//示例:
//
//给定 matrix = [
//[3, 0, 1, 4, 2],
//[5, 6, 3, 2, 1],
//[1, 2, 0, 1, 5],
//[4, 1, 0, 1, 7],
//[1, 0, 3, 0, 5]
//]
//
//sumRegion(2, 1, 4, 3) -> 8
//sumRegion(1, 1, 2, 2) -> 11
//sumRegion(1, 2, 2, 4) -> 12
//说明:
//
//你可以假设矩阵不可变。
//会多次调用 sumRegion 方法。
//你可以假设 row1 ≤ row2 且 col1 ≤ col2。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/range-sum-query-2d-immutable
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type NumMatrix struct {
	matrixNum [][]int
}

func ConstructorMatrix(matrix [][]int) NumMatrix {
	return NumMatrix{matrixNum: matrix}
}

//func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
//	var sum int
//	for i := row1; i <= row2; i++ {
//		for j := col1; j <= col2; j++ {
//			sum += this.matrixNum[i][j]
//		}
//	}
//	return sum
//}
var cacheMatrix = make([][]int,100,0)
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return cacheMatrix[row2+1][col2+1] - cacheMatrix[row2+1][col1] - cacheMatrix[row1][col2+1] + cacheMatrix[row1][col1]
}
func CacheSumMatrix(matrix [][]int) () {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	for i := 0; i < len(matrix)+1; i++ {
		for j := 0; j < len(matrix[0])+1; j++ {
			cacheMatrix[i+1][j+1] = cacheMatrix[i+1][j] + cacheMatrix[i][j+1] + matrix[i][j] - cacheMatrix[i][j]
		}
	}
}

//private int[][] dp
//
//public NumMatrix(int[][] matrix) {
//if (matrix.length == 0 || matrix[0].length == 0) return;
//dp = new int[matrix.length + 1][matrix[0].length + 1];
//for (int r = 0; r < matrix.length; r++) {
//for (int c = 0; c < matrix[0].length; c++) {
//dp[r + 1][c + 1] = dp[r + 1][c] + dp[r][c + 1] + matrix[r][c] - dp[r][c];
//}
//}
//}
//
//public int sumRegion(int row1, int col1, int row2, int col2) {
//return dp[row2 + 1][col2 + 1] - dp[row1][col2 + 1] - dp[row2 + 1][col1] + dp[row1][col1];
//}
