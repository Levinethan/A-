package main

import "math"

//A星算法  地图上的点结构
type AstarPoint struct {
	Point
	farther *AstarPoint

	gVal int
	hVal int
	fVal int

	//A星算法结构

	//使用优先队列
}

func NewAstarPoint (p *Point,father *AstarPoint , end *AstarPoint) (ap *AstarPoint){
	ap = &AstarPoint{
		Point:   *p,
		farther: father,
		gVal:    0,
		hVal:    0,
		fVal:    0,
	}

	if end!= nil{
		ap.CalcFval(end)
	}
	return ap
}

func (this *AstarPoint)CalcGval () int{
	//gval 表示从初始节点到任意节点n的代价
	if this.farther!=nil{
		deltaX := math.Abs(float64(this.farther.x-this.x))
		deltaY := math.Abs(float64(this.farther.y-this.y))
		if deltaX==1 && deltaY==0{
			//移动了一步
			this.gVal = this.farther.gVal +10

		}else if deltaX==0 && deltaY==1{
			this.gVal = this.farther.gVal+10
		}else if deltaY==1 && deltaX==1{
			this.gVal = this.farther.gVal+14
		}else {
			panic("error")
		}
	}

	return this.gVal
	
}

func (this *AstarPoint)CalcHval(end *AstarPoint)int {
	this.hVal =  int(math.Abs(float64(end.x-this.x))+math.Abs(float64(end.y-this.y)))
	//计算当前节点与目标节点的差值

	return this.hVal
}
func (this *AstarPoint)CalcFval(end *AstarPoint)int {
	this.fVal= this.CalcGval() + this.CalcHval(end)
	return this.fVal
}