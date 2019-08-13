package tensor

//Int32
type CalculatorInt interface {
	AddInt(i *Int32) Int32    //加法
	SubInt(i *Int32) Int32    //减法
	DivInt(i *Int32) Int32    //除法
	MulInt(i *Int32) Int32    //乘法
	MatMulInt(i *Int32) Int32 //矩阵乘法
}

type DataChan struct {
	index int
	data  interface{}
}
type DataChans []DataChan

func (t *DataChans) Len() int {
	return len(*t)
}
func (t *DataChans) Less(i, j int) bool {

}
func (t *DataChans) Swap(i, j int) {
}

func (t *Int32) AddInt(i *Int32) (rest Int32) {
	rest.tensorInt32 = make([]int32, 0)
	//当t是标量
	if t.Size() == 1 {

		for _, iV := range i.tensorInt32 {
			rest.tensorInt32 = append(rest.tensorInt32, t.tensorInt32[0]+iV)
			rest.Shape = i.Shape
			return
		}

	}
	//当i是标量
	if i.Size() == 1 {
		for _, tV := range t.tensorInt32 {
			rest.tensorInt32 = append(rest.tensorInt32, i.tensorInt32[0]+tV)
			rest.Shape = t.Shape
			return
		}
	}
	//形状相等
	if i.GetShape().Equ(t.GetShape()) {
		for j, _ := range t.tensorInt32 {
			rest.tensorInt32 = append(rest.tensorInt32, t.tensorInt32[j]+i.tensorInt32[j])
		}
	}

	/*//申请通道[2]int32
	restChan := make(chan DataChan, len(t.tensorInt32))
	for j, tV := range t.tensorInt32 {
		go func(j int) {
			var temp int32
			for _, iV := range i.tensorInt32 {
				temp += tV + iV
			}
			restChan <- DataChan{j, temp}
		}(j)

	}

	var tempChan DataChans
	for range t.tensorInt32 {
		tempChan = append(tempChan, <-restChan)
	}
	sort.Sort(&tempChan)
	for _, v := range tempChan {

	}*/
}
func ProdIntsGo(a []int, num int) (retVal int) {
	aLen := len(a)
	if aLen == 0 {
		return
	}
	n := num
	if num > aLen {
		n = aLen
	}

	//fmt.Println(n)
	restChan := make(chan int, n)
	m := aLen / n
	for i := 0; i < n; i++ {
		go func(i int) {
			var t = 1                //临时变量
			ints := a[i*m : (i+1)*m] // i < n-1时

			if i == n-1 {
				ints = a[(n-1)*m:]
			}

			for _, v := range ints {
				t *= v
			}
			restChan <- t
		}(i)
	}

	var total = 1
	for i := 0; i < n; i++ {
		total *= <-restChan
	}
	retVal = total

	return
}
func (t Int32) SubInt(i Int32) Int32 {
	panic("implement me")
}

func (t Int32) DivInt(i Int32) Int32 {
	panic("implement me")
}

func (t Int32) MulInt(i Int32) Int32 {
	panic("implement me")
}

func (t Int32) MatMulInt(i Int32) Int32 {
	panic("implement me")
}
