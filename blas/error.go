package blas

const (
	ZeroIncX = "blas: zero x index increment"
	ZeroIncY = "blas: zero y index increment"
	MLT0     = "blas: m < 0"
	NLT0     = "blas: n < 0"
	KLT0     = "blas: k < 0"
	KLLT0    = "blas: kL < 0"
	KULT0    = "blas: kU < 0"

	BadUplo      = "blas: illegal triangle"
	BadTranspose = "blas: illegal transpose"
	BadDiag      = "blas: illegal diagonal"
	BadSide      = "blas: illegal side"
	BadFlag      = "blas: illegal rotm flag"

	BadLdA = "blas: bad leading dimension of A"
	BadLdB = "blas: bad leading dimension of B"
	BadLdC = "blas: bad leading dimension of C"

	ShortX  = "blas: insufficient length of x"
	ShortY  = "blas: insufficient length of y"
	ShortAP = "blas: insufficient length of ap"
	ShortA  = "blas: insufficient length of a"
	ShortB  = "blas: insufficient length of b"
	ShortC  = "blas: insufficient length of c"

	BadLength = "vector length mismatch"
	NegInc    = "negative vector increment"
)
