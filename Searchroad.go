package main

import "container/heap"

type SearchRoad struct {

	theMap *Map  //地图

	start AstarPoint //开始
	end AstarPoint//结束

	closeLi map[string]*AstarPoint  //关闭 不通的路


	openLi OpenList

	openSet map[string]*AstarPoint
	//set 去掉重复

	TheRoad []*AstarPoint
}

func NewSearchRoad(startx,starty,endx,endy int , m *Map) *SearchRoad  {
	sr := &SearchRoad{}
	sr.theMap=m
	sr.start=*NewAstarPoint(&Point{startx,starty,"S"},nil,nil)
	//开始节点
	sr.end=*NewAstarPoint(&Point{endx,endy,"E"},nil,nil)
	sr.TheRoad = make([]*AstarPoint,0)  //开辟内存  存储路
	sr.openSet = make(map[string]*AstarPoint,m.MaxX+m.MaxY)  //开放集合
	sr.closeLi = make(map[string]*AstarPoint,m.MaxX+m.MaxY)
	heap.Init(&sr.openLi)
	//初始化栈
	heap.Push(&sr.openLi,&sr.start)
	//压入开始节点
	sr.openSet[PointAsKey(sr.start.x,sr.start.y)] = &sr.start
	for k,v := range m.blocks{
		sr.closeLi[k]= NewAstarPoint(v,nil,nil)
		//所有的障碍加入block

	}
	return sr
}
//A星算法核心
func (this *SearchRoad)FindoutShortestPath() bool  {
	for len(this.openLi) >0 {
		//如果开放节点大于零 永远循环下去
		//如果找不到路 从开放节点中取出  放入关闭节点
		x := heap.Pop(&this.openLi) //取出一个节点
		curPoint := x.(*AstarPoint) //取得当前节点
		delete(this.openSet,PointAsKey(curPoint.x,curPoint.y)) //删除开放列表
		this.closeLi[PointAsKey(curPoint.x,curPoint.y)] = curPoint //障碍走过的路 加入关闭列表
		adjacs := this.theMap.GetAdjaoentPoint(&curPoint.Point) //取出所有邻居节点
		for _ , p := range adjacs{
			thsAp := NewAstarPoint(p,curPoint,&this.end) //创建A星节点
			if PointAsKey(thsAp.x,thsAp.y) == PointAsKey(this.end.x,this.end.y){
				//找到节点  然后标记
				for thsAp.farther !=nil{
					this.TheRoad = append(this.TheRoad,thsAp)
					thsAp.view = "*"
					thsAp= thsAp.farther //返回上一个节点
				}
				return true
			}
			_ , ok := this.closeLi[PointAsKey(p.x,p.y)]
			if ok{
				continue
			}

			exitAP , ok := this.openSet[PointAsKey(p.x,p.y)]  //取出节点
			if !ok{
				//如果存在
				heap.Push(&this.openLi,thsAp)  //节点不存在 压入
				this.openSet[PointAsKey(thsAp.x,thsAp.y)] = thsAp //放入开放列表
			}else {
				oldval , oldfather  :=exitAP.gVal,exitAP.farther
				exitAP.farther = curPoint //当前的父亲节点
				exitAP.CalcGval() //计算最短的值
				if exitAP.gVal > oldval{
					//新的节点距离 比老节点还要短
					exitAP.farther = oldfather
					exitAP.gVal = oldval
				}


			}
		}
	}
	return false
}