package bincode

func LeftRightBounds(rb, re int, vlen int) (ii, ij int) {
	// границы как в python:
	// положительный - имеет максимум до длины (len)
	// отрицательный - считается с конца с минимумом -длина
	// если выходит за макс. границу - возвращаем пустой слайс
	// если выходит за мин. границу - считаем =0

	// правая граница как в python - исключается

	// левая граница включая
	ii = rb

	switch {
	case ii > 0:
		if ii >= vlen {
			ii = vlen - 1
		}
	case ii < 0:
		ii += vlen
		if ii < 0 {
			ii = 0
		}
	}

	ij = re

	switch {
	case ij > 0:
		if ij > vlen {
			ij = vlen
		}
	case ij < 0:
		ij += vlen
		if ij < 0 {
			ij = 0
		}
	}
	return
}

