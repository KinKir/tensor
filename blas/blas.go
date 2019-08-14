package blas

type BLAS struct{}

// Uplo specifies whether a matrix is upper or lower triangular.
// Uplo指定矩阵是上三角形还是下三角形。
type Uplo byte

const (
	Upper Uplo = 'U'
	Lower Uplo = 'L'
	All   Uplo = 'A'
)

// Diag specifies whether a matrix is unit triangular.
// Diag指定矩阵是否为单位三角形
type Diag byte

const (
	NonUnit Diag = 'N'
	Unit    Diag = 'U'
)
