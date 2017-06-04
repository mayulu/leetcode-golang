package leetcode_go_mayulu

func Convert(s string, numRows int) string {
	sl := len(s)
	if numRows <= 1 || sl <= numRows {
		return s
	}

	blockCols := int(math.Ceil(float64(sl) / float64(numRows)))

	numInBlock := 2*numRows - 2

	index1 := 0
	index2 := 0

	resultBuff := []byte{}

	for iR := 1; iR <= numRows; iR++ {
		for jBC := 1; jBC <= blockCols; jBC++ {
			index1 = numInBlock*(jBC-1) + iR
			if index1 <= sl {
				resultBuff = append(resultBuff, s[index1-1])
			}

			if iR != 1 && iR != numRows {
				index2 = numInBlock*(jBC-1) + 2*numRows - iR
				if index2 <= sl {
					resultBuff = append(resultBuff, s[index2-1])
				}

			}

		}

	}

	return string(resultBuff)

}
