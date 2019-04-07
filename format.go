package main

import "fmt"

//按照numpy的格式输出
func (m *Matrix64) String() string {
	limit := 8
	var str = "["
	countRow := 0
	for i := 0; i < m.row; i++ {
		if m.row > limit { //行数大于limit，只显示前后3行
			if i > 2 && i < (m.row-3) {
				if countRow == 0 {
					str += "...\n"
					countRow++
				}
				continue
			}
		}
		if i != 0 {
			str += " "
		}
		str += "["
		countCol := 0
		for j := 0; j < m.col; j++ {
			if m.col > limit { //列数大于limit，只显示前后3列
				if j > 2 && j < (m.col-3) {
					if countCol == 0 {
						str += "  \t...\t"
						countCol++
					}
					continue
				}
			}
			if j != 0 {
				str += "  "
			}
			str += fmt.Sprintf("%f", m.mat[m.col*i+j])
		}
		str += "]"
		if i != m.row-1 {
			str += "\n"
		}
	}
	str += "]"
	return str
}

func (m *Matrix32) String() string {
	limit := 8
	var str = "["
	countRow := 0
	for i := 0; i < m.row; i++ {
		if m.row > limit { //行数大于limit，只显示前后3行
			if i > 2 && i < (m.row-3) {
				if countRow == 0 {
					str += "...\n"
					countRow++
				}
				continue
			}
		}
		if i != 0 {
			str += " "
		}
		str += "["
		countCol := 0
		for j := 0; j < m.col; j++ {
			if m.col > limit { //列数大于limit，只显示前后3列
				if j > 2 && j < (m.col-3) {
					if countCol == 0 {
						str += "  \t...\t"
						countCol++
					}
					continue
				}
			}
			if j != 0 {
				str += "  "
			}
			str += fmt.Sprintf("%f", m.mat[m.col*i+j])
		}
		str += "]"
		if i != m.row-1 {
			str += "\n"
		}
	}
	str += "]"
	return str
}

/*
func (m *Matrix64) String() string {
	limit := 7
	var str = "["
	switch {
	case m.row <= limit && m.col <= limit:
		for i := 0; i < m.row; i++ {
			if i != 0 {
				str += " "
			}
			str += "["
			for j := 0; j < m.col; j++ {
				if j != 0 {
					str += "  "
				}
				str += fmt.Sprintf("%f", m.mat[m.col*i+j])
			}
			str += "]"
			if i != m.row-1 {
				str += "\n"
			}
		}
	case m.row > limit && m.col <= limit:
		countRow := 0
		for i := 0; i < m.row; i++ {
			if i > 2 && i < (m.row-3) {
				if countRow == 0 {
					str += "...\n"
					countRow++
				}
				continue
			}
			if i != 0 {
				str += " "
			}
			str += "["
			for j := 0; j < m.col; j++ {
				if j != 0 {
					str += "  "
				}
				str += fmt.Sprintf("%f", m.mat[m.col*i+j])
			}
			str += "]"
			if i != m.row-1 {
				str += "\n"
			}
		}
	case m.row <= limit && m.col > limit:
		for i := 0; i < m.row; i++ {
			if i != 0 {
				str += " "
			}
			str += "["
			countCol := 0
			for j := 0; j < m.col; j++ {
				if j > 2 && j < (m.col-3) {
					if countCol == 0 {
						str += "  \t...\t"
						countCol++
					}
					continue
				}
				if j != 0 {
					str += "  "
				}
				str += fmt.Sprintf("%f", m.mat[m.col*i+j])
			}
			str += "]"
			if i != m.row-1 {
				str += "\n"
			}
		}
	default:
		countRow := 0
		for i := 0; i < m.row; i++ {
			if i > 2 && i < (m.row-3) {
				if countRow == 0 {
					str += "...\n"
					countRow++
				}
				continue
			}
			if i != 0 {
				str += " "
			}
			str += "["
			countCol := 0
			for j := 0; j < m.col; j++ {
				if j > 2 && j < (m.col-3) {
					if countCol == 0 {
						str += "  \t...\t"
						countCol++
					}
					continue
				}
				if j != 0 {
					str += "  "
				}
				str += fmt.Sprintf("%f", m.mat[m.col*i+j])
			}
			str += "]"
			if i != m.row-1 {
				str += "\n"
			}
		}
	}
	str += "]"
	return str
}
*/
