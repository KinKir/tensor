package f32

func L1Norm(x []float32) (sum float32) {
	for _, v := range x {
		sum += abs(v)
	}
	return sum
}
func abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
func Abs(x []float32) []float32 {
	for i, v := range x {
		if v > 0 {
			x[i] = v
		} else {
			x[i] = -v
		}
	}
	return x
}

// Div is
//  for i, v := range s {
//  	dst[i] /= v
//  }
func Div(dst, s []float32) {
	for i, v := range s {
		dst[i] /= v
	}
}

// L1NormInc is
//  for i := 0; i < n*incX; i += incX {
//  	sum += math.Abs(x[i])
//  }
//  return sum
func L1NormInc(x []float32, n, incX int) (sum float32) {
	for i := 0; i < n*incX; i += incX {
		sum += abs(x[i])
	}
	return sum
}

// AddConst is
//  for i := range x {
//  	x[i] += alpha
//  }
func AddConst(alpha float32, x []float32) []float32 {
	for i := range x {
		x[i] += alpha
	}
	return x
}

// Add is
//  for i, v := range s {
//  	dst[i] += v
//  }
func Add(dst, s []float32) []float32 {
	for i, v := range s {
		dst[i] += v
	}
	return dst
}

// AxpyUnitary is
//  for i, v := range x {
//  	y[i] += alpha * v
//  }
func AxpyUnitary(alpha float32, x, y []float32) []float32 {
	for i, v := range x {
		y[i] += alpha * v
	}
	return y

}

// AxpyUnitaryTo is
//  for i, v := range x {
//  	dst[i] = alpha*v + y[i]
//  }
func AxpyUnitaryTo(dst []float32, alpha float32, x, y []float32) []float32 {
	for i, v := range x {
		dst[i] = alpha*v + y[i]
	}
	return dst
}

// AxpyInc is
//  for i := 0; i < int(n); i++ {
//  	y[iy] += alpha * x[ix]
//  	ix += incX
//  	iy += incY
//  }
// NOTE: ？？？？？
func AxpyInc(alpha float32, x, y []float32, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		y[iy] += alpha * x[ix]
		ix += incX
		iy += incY
	}
}

// AxpyIncTo is
//  for i := 0; i < int(n); i++ {
//  	dst[idst] = alpha*x[ix] + y[iy]
//  	ix += incX
//  	iy += incY
//  	idst += incDst
//  }
func AxpyIncTo(dst []float32, incDst, idst uintptr, alpha float32, x, y []float32, n, incX, incY, ix, iy uintptr) {

}

// CumSum is
//  if len(s) == 0 {
//  	return dst
//  }
//  dst[0] = s[0]
//  for i, v := range s[1:] {
//  	dst[i+1] = dst[i] + v
//  }
//  return dst
func CumSum(dst, s []float32) []float32 {

}

// CumProd is
//  if len(s) == 0 {
//  	return dst
//  }
//  dst[0] = s[0]
//  for i, v := range s[1:] {
//  	dst[i+1] = dst[i] * v
//  }
//  return dst
func CumProd(dst, s []float32) []float32 {

}

// DivTo is
//  for i, v := range s {
//  	dst[i] = v / t[i]
//  }
//  return dst
func DivTo(dst, x, y []float32) []float32 {

}

// DotUnitary is
//  for i, v := range x {
//  	sum += y[i] * v
//  }
//  return sum
func DotUnitary(x, y []float32) (sum float32) {

}

// DotInc is
//  for i := 0; i < int(n); i++ {
//  	sum += y[iy] * x[ix]
//  	ix += incX
//  	iy += incY
//  }
//  return sum
func DotInc(x, y []float32, n, incX, incY, ix, iy uintptr) (sum float32) {

}

// L1Dist is
//  var norm float32
//  for i, v := range s {
//  	norm += math.Abs(t[i] - v)
//  }
//  return norm
func L1Dist(s, t []float32) float32 {

}

// LinfDist is
//  var norm float32
//  if len(s) == 0 {
//  	return 0
//  }
//  norm = math.Abs(t[0] - s[0])
//  for i, v := range s[1:] {
//  	absDiff := math.Abs(t[i+1] - v)
//  	if absDiff > norm || math.IsNaN(norm) {
//  		norm = absDiff
//  	}
//  }
//  return norm
func LinfDist(s, t []float32) float32 {

}

// ScalUnitary is
//  for i := range x {
//  	x[i] *= alpha
//  }
func ScalUnitary(alpha float32, x []float32) {

}

// ScalUnitaryTo is
//  for i, v := range x {
//  	dst[i] = alpha * v
//  }
func ScalUnitaryTo(dst []float32, alpha float32, x []float32) {

}

// ScalInc is
//  var ix uintptr
//  for i := 0; i < int(n); i++ {
//  	x[ix] *= alpha
//  	ix += incX
//  }
func ScalInc(alpha float32, x []float32, n, incX uintptr) {

}

// ScalIncTo is
//  var idst, ix uintptr
//  for i := 0; i < int(n); i++ {
//  	dst[idst] = alpha * x[ix]
//  	ix += incX
//  	idst += incDst
//  }
func ScalIncTo(dst []float32, incDst uintptr, alpha float32, x []float32, n, incX uintptr) {

}

// Sum is
//  var sum float32
//  for i := range x {
//      sum += x[i]
//  }
func Sum(x []float32) float32 {

}
