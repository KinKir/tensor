package tensor

import (
	"fmt"
	"reflect"
)

type Shape []int

type Shaper interface {
	Size() int
	Dims() int
	IsMatrix() bool
	IsRowVec() bool
	IsColVec() bool
	IsVector() bool
	IsScalarEquiv() bool
}

func (s Shape) Size() (size int) {
	if len(s) == 0 {
		return
	}
	size = 1
	for _, v := range s {
		size *= v
	}
	return
}

//是否是标量,列如：2,[2]
func (s Shape) IsScalar() bool {

	return len(s) == 0 || (len(s) == 1 && s[0] == 1)
}

//维度数
func (s Shape) Dims() int { return len(s) }

//是否是一个矩阵
func (s Shape) IsMatrix() bool { return len(s) == 2 }

//是否是行向量
func (s Shape) IsRowVec() bool { return len(s) == 2 && (s[0] == 1 && s[1] > 1) }

//是否是列向量，例如[2,1]
func (s Shape) IsColVec() bool { return len(s) == 2 && (s[1] == 1 && s[0] > 1) }

//是否是一个向量
func (s Shape) IsVector() bool { return s.IsColVec() || s.IsRowVec() || (len(s) == 1 && s[0] > 1) }

func (s Shape) IsScalarEquiv() bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if v != 1 { //如果长度均为1，则是一个类似标量的张量
			return false
		}
	}
	return true
}
func (s Shape) Equ(shape Shape) bool  {
	return  reflect.DeepEqual(s,shape)
}

//格式化输出
func (s Shape) Format(st fmt.State, r rune) {
	switch r {
	case 'v', 's':
		_, _ = st.Write([]byte("("))
		for i, v := range s {
			_, _ = fmt.Fprintf(st, "%d", v)
			if i < len(s)-1 {
				_, _ = st.Write([]byte(", "))
			}
		}
		_, _ = st.Write([]byte(")"))
	default:
		_, _ = fmt.Fprintf(st, "%v", []int(s))
	}
}
