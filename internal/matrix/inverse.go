package matrix

import (
	"fmt"
)

func (m Matrix) Inverse() Matrix {
	if m.row_size != 4 || m.col_size != 4 {
		panic(fmt.Errorf("Inverse calculation only implemented for 4x4 matrices, called with: \n%s", m.String()))
	}

	det := determinant(m)
	if det == 0 {
		panic(fmt.Errorf("non-invertible matrix, determinant is zero for \n%s", m.String()))
	}

	inv := Matrix{
		data:     [16]float64{},
		row_size: 4,
		col_size: 4,
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sign := 1.0
			if (i+j)%2 != 0 {
				sign = -1.0
			}
			inv.data[j*inv.col_size+i] = sign * subDet(m, i, j) / det
		}
	}

	return inv
}

// determinant for 4x4 matrix
func determinant(m Matrix) float64 {
	det := 0.0
	for j := 0; j < 4; j++ {
		if j%2 == 0 {
			det += m.At(0, j) * subDet(m, 0, j)
		} else {
			det -= m.At(0, j) * subDet(m, 0, j)
		}
	}
	return det
}

// determinant for a 3x3 matrix by excluding row, and col from the original 4x4 matrix
func subDet(m Matrix, row, col int) float64 {
	sub := [9]float64{}
	index := 0
	for i := 0; i < 4; i++ {
		if i == row {
			continue
		}
		for j := 0; j < 4; j++ {
			if j == col {
				continue
			}
			sub[index] = m.At(i, j)
			index++
		}
	}
	// det(A)=a(ei−fh)−b(di−fg)+c(dh−eg)
	// {
	// 	a, b, c,
	// 	d, e, f,
	// 	g, h, i,
	// }
	return sub[0]*(sub[4]*sub[8]-sub[5]*sub[7]) - sub[1]*(sub[3]*sub[8]-sub[5]*sub[6]) + sub[2]*(sub[3]*sub[7]-sub[4]*sub[6])
}
