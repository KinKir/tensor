package f32

import (
	"github.com/chenquan/tensor/blas"
)

//实现单精度浮点类型的基础线性代数程序集

// Vector represents a vector with an associated element increment.
// Vector表示具有关联元素增量的向量。
type Vector struct {
	N    int
	Inc  int
	Data []float32
}

// General represents a matrix using the conventional storage scheme.
// General表示使用传统存储方案的矩阵。
type General struct {
	Rows, Cols int
	Stride     int
	Data       []float32
}

// Band represents a band matrix using the band storage scheme.
// Band表示使用频带存储方案的频段矩阵。
type Band struct {
	Rows, Cols int
	KL, KU     int
	Stride     int
	Data       []float32
}

// Triangular represents a triangular matrix using the conventional storage scheme.
//三角形表示使用传统存储方案的三角形矩阵。
type Triangular struct {
	N      int
	Stride int
	Data   []float32
	Uplo   blas.Uplo
	Diag   blas.Diag
}

// TriangularPacked represents a triangular matrix using the packed storage scheme.
// TriangularPacked使用打包存储方案表示三角矩阵。
type TriangularPacked struct {
	N    int
	Data []float32
	Uplo blas.Uplo
	Diag blas.Diag
}

// Symmetric represents a symmetric matrix using the conventional storage scheme.
//对称表示使用传统存储方案的对称矩阵。
type Symmetric struct {
	N      int
	Stride int
	Data   []float32
	Uplo   blas.Uplo
}

// SymmetricBand represents a symmetric matrix using the band storage scheme.
// SymmetricBand 使用频段存储方案表示对称矩阵。
type SymmetricBand struct {
	N, K   int
	Stride int
	Data   []float32
	Uplo   blas.Uplo
}

// SymmetricPacked represents a symmetric matrix using the packed storage scheme.
// SymmetricPacked使用打包存储方案表示对称矩阵。
type SymmetricPacked struct {
	N    int
	Data []float32
	Uplo blas.Uplo
}

func Dot(x, y Vector) float32 {
	if x.N != y.N {
		panic(blas.BadLength)
	}
	if x.Inc == 0 {
		panic(blas.ZeroIncX)
	}
	if y.Inc == 0 {
		panic(blas.ZeroIncY)
	}
	n := x.N
	if n <= 0 {
		if n == 0 {
			return 0
		}
		panic(blas.NLT0)
	}

	if x.Inc == 1 && y.Inc == 1 {
		if len(x.Data) < n {
			panic(blas.ShortX)
		}
		if len(y.Data) < n {
			panic(blas.ShortY)
		}
		return DotUnitary(x.Data[:n], y.Data[:n])
	}
	var ix, iy int
	if x.Inc < 0 {
		ix = (-n + 1) * x.Inc
	}
	if y.Inc < 0 {
		iy = (-n + 1) * y.Inc
	}
	if ix >= len(x.Data) || ix+(n-1)*x.Inc >= len(x.Data) {
		panic(blas.ShortX)
	}
	if iy >= len(y.Data) || iy+(n-1)*y.Inc >= len(y.Data) {
		panic(blas.ShortY)
	}
	return DotInc(x.Data, y.Data, uintptr(n), uintptr(x.Inc), uintptr(y.Inc), uintptr(ix), uintptr(iy))

}
