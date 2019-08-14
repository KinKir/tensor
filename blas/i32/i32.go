package i32

import "github.com/chenquan/tensor/blas"

//实现32位整型类型的基础线性代数程序集

// Vector represents a vector with an associated element increment.
// Vector表示具有关联元素增量的向量。
type Vector struct {
	N int
	Inc  int
	Data []int32
}

// General represents a matrix using the conventional storage scheme.
// General表示使用传统存储方案的矩阵。
type General struct {
	Rows, Cols int
	Stride     int
	Data       []int32
}

// Band represents a band matrix using the band storage scheme.
// Band表示使用频带存储方案的频段矩阵。
type Band struct {
	Rows, Cols int
	KL, KU     int
	Stride     int
	Data       []int32
}

// Triangular represents a triangular matrix using the conventional storage scheme.
//三角形表示使用传统存储方案的三角形矩阵。
type Triangular struct {
	N      int
	Stride int
	Data   []int32
	Uplo   blas.Uplo
	Diag   blas.Diag
}

// TriangularPacked represents a triangular matrix using the packed storage scheme.
// TriangularPacked使用打包存储方案表示三角矩阵。
type TriangularPacked struct {
	N    int
	Data []int32
	Uplo blas.Uplo
	Diag blas.Diag
}

// Symmetric represents a symmetric matrix using the conventional storage scheme.
//对称表示使用传统存储方案的对称矩阵。
type Symmetric struct {
	N      int
	Stride int
	Data   []int32
	Uplo   blas.Uplo
}

// SymmetricBand represents a symmetric matrix using the band storage scheme.
// SymmetricBand 使用频段存储方案表示对称矩阵。
type SymmetricBand struct {
	N, K   int
	Stride int
	Data   []int32
	Uplo   blas.Uplo
}

// SymmetricPacked represents a symmetric matrix using the packed storage scheme.
// SymmetricPacked使用打包存储方案表示对称矩阵。
type SymmetricPacked struct {
	N    int
	Data []int32
	Uplo blas.Uplo
}
