package p003x

/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1:

输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true

示例 2:

输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。

说明:

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
给定数独序列只包含数字 1-9 和字符 '.' 。
给定数独永远是 9x9 形式的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-sudoku
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isValidSudoku(board [][]byte) bool {
    var checkRow, checkCol, checkBlk int
    var ch byte
    for i := 0; i < 9; i++ {
        checkRow, checkCol, checkBlk = 0, 0, 0

        for j := 0; j < 9; j++ {
            // 检查第i行
            ch = board[i][j]
            if ch != '.' {
                if checkRow&(1<<(ch-'0')) != 0 {
                    return false
                }
                checkRow |= (1 << (ch - 48))
            }

            // 检查第i列
            ch = board[j][i]
            if ch != '.' {
                if checkCol&(1<<(ch-'0')) != 0 {
                    return false
                }
                checkCol |= (1 << (ch - 48))
            }

            // 第i块 坐标偏移量(i%3)*3, (i/3)*3
            // 块内第j格  坐标偏移量j%3, j/3
            ch = board[(i%3)*3+j%3][(i/3)*3+j/3]
            if ch != '.' {
                if checkBlk&(1<<(ch-'0')) != 0 {
                    return false
                }
                checkBlk |= (1 << (ch - 48))
            }
        }
    }
    return true
}
