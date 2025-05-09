package test

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		if !isSudoKu(board, len(board[i]) - 1, i, 0, i) {
			return false
		}
	}

	for i := 0; i < len(board[0]); i++ {
		if !isSudoKu(board, i, len(board) -1, i, 0) {
			return false
		}
	}

	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board[i]); j+= 3 {
			if !isSudoKu(board, j + 2, i + 2, j, i) {
				return false
			}
		}
	}

	return true
}

func isSudoKu(board [][]byte, rightX, rightY, leftX, leftY int) bool {
	type void struct {}		
	set := make(map[byte]void)
	for y := leftY; y <= rightY; y++ {
		for x := leftX; x <= rightX; x++ {
			if board[y][x] == '.' {
				continue
			}
			_, ok := set[board[y][x]]
			if ok {
				return false
			}
			set[board[y][x]] = void{}
		}
	}
	return true
}