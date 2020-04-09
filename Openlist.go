package main

type OpenList [] *AstarPoint

//openlist  起到栈的作用


func (self OpenList)Len() int {
	return len(self)

}

func (self OpenList)Less(i,j int) bool {
	return self[i].fVal < self[j].fVal
}

func  (self OpenList)Swap(i,j int){
	self[i],self[j]=self[j],self[i]
}

func (this *OpenList)Push (data interface{}){

	*this = append (*this , data.(*AstarPoint))  //节点加入栈中
}

func (this *OpenList)Pop() interface{}{

	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0:n-1]
	return x
}
