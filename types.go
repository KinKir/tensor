package tensor

import "reflect"

//整型
//type tensorInt8 []int8
//type tensorInt16 []int16
//type tensorInt []int
type tensorInt32 []int32
type tensorInt64 []int64

//
//type Int8 struct {
//	tensorInt8
//	Shape
//}
//type Int16 struct {
//	tensorInt16
//	Shape
//}
//type Int struct {
//	tensorInt
//	Shape
//}
type Int32 struct {
	tensorInt32
	Shape
}
type Int64 struct {
	tensorInt64
	Shape
}

//浮点型
type tensorFloat32 []float32
type tensorFloat64 []float64
type Float32 struct {
	tensorFloat32
	Shape
}
type Float64 struct {
	tensorFloat64
	Shape
}
type Typer interface {
	Size() int
	Type() reflect.Type
	GetShape() Shape
}

/*//Int8
func (i *Int8) Size() int {
	return len(i.shape)
}
func (i *Int8) Type() reflect.Type {
	return reflect.TypeOf(i)
}
func (i *Int8) GetShape() Shape {
	return i.Shape
}

//Int16
func (i *Int16) Size() int {
	return len(i.shape)
}
func (i *Int16) Type() reflect.Type {
	return reflect.TypeOf(i)
}
func (i *Int16) GetShape() Shape {
	return i.Shape
}

//Int
func (i *Int) Size() int {
	return len(i.shape)
}
func (i *Int) Type() reflect.Type {
	return reflect.TypeOf(i)
}
func (i *Int) GetShape() Shape {
	return i.Shape
}*/

//Int32
func (t *Int32) Size() int {
	return len(t.tensorInt32)
}
func (t *Int32) Type() reflect.Type {
	return reflect.TypeOf(t)
}
func (t *Int32) GetShape() Shape {
	return t.Shape
}

//Float32
func (f *Float32) Size() int {
	return len(f.Shape)
}
func (f *Float32) Type() reflect.Type {
	return reflect.TypeOf(f)
}
func (f *Float32) GetShape() Shape {
	return f.Shape
}

//Float64
func (f *Float64) Size() int {
	return len(f.Shape)
}
func (f *Float64) Type() reflect.Type {
	return reflect.TypeOf(f)
}
func (f *Float64) GetShape() Shape {
	return f.Shape
}
