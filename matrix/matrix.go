package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct{
	arr [][]int
	rows, cols int
}

func get2DSlice(rows, cols int) [][]int {		// Just a utility
	arr := make([][]int, rows)
	for i:=0; i<rows; i++ {
		arr[i] = make([]int, cols)
	}
	return arr
}

func (m *Matrix) Rows() [][]int {
	rows := get2DSlice(m.rows, m.cols)

	for r := 0; r<m.rows; r++ {
		for c := 0; c<m.cols; c++ {
			rows[r][c] = m.arr[r][c]
		}
	}

	return rows
}

func (m *Matrix) Cols() [][]int {
	cols := get2DSlice(m.cols, m.rows)
	
	for c := 0; c<m.cols; c++ {
		for r := 0; r<m.rows; r++ {
			cols[c][r] = m.arr[r][c]
		}
	}

	return cols
}

func (m *Matrix) Set(r, c, val int) bool {
	if r<m.rows && c<m.cols && r>=0 && c>=0 {
		m.arr[r][c] = val
		return true	
	}
	return false
}

func New(s string) (*Matrix, error) {
	lines := strings.Split(s, "\n")

	mat := Matrix{}
	mat.rows = len(lines)
	mat.arr = make([][]int, mat.rows)


	for idx, line := range lines {
		line = strings.Trim(line, " ")
		cells := strings.Split(line, " ")
		
		cellNum := len(cells)
		if idx == 0 {
			mat.cols = cellNum
		} else if (cellNum != mat.cols) {
			return nil, errors.New("uneven cols")
		}

		for _, cell := range cells {
			cellInt, err := strconv.Atoi(cell)
			if err != nil {
				return nil, errors.New("invalid cell")
			}
			mat.arr[idx] = append(mat.arr[idx], cellInt)
		}
	}

	return &mat, nil
}

