package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Map struct {

	points [][] Point

	blocks map[string]*Point

	MaxX int
	MaxY int
}

func PointAsKey(x,y int) (key string){
	key = strconv.Itoa(x) + "," + strconv.Itoa(y)
	return key  //坐标转化为字符串
}

func NewMap (charMap []string) (m Map){
	m.points = make([][]Point,len(charMap))

	m.blocks = make(map[string]*Point,len(charMap)*2)  //储存两倍边长

	for x, row := range charMap {
		cols := strings.Split(row," ") //基于空格切割
		m.points[x] = make([]Point , len(cols)) //二维数组每个元素开辟内存
		for y, view := range cols {

			m.points [x][y] = Point{x,y,view} //* - .
			if view == "X"{
				//标记障碍
				m.blocks[PointAsKey(x,y)] = &m.points[x][y]
			}
		}
	}
	m.MaxX = len(m.points)
	m.MaxY = len(m.points[0])
	return m

}


//抓取相邻节点  返回一个集合
func (this *Map)GetAdjaoentPoint(curPoint *Point) (Adjaoent[]* Point){
	//寻路算法设置为八个方向
	if x,y := curPoint.x , curPoint.y-1 ;x>=0 && x <this.MaxX && y>=0 && y <this.MaxY{
			Adjaoent = append(Adjaoent,&this.points[x][y])
	}
	if x,y := curPoint.x+1 , curPoint.y-1 ;x>=0 && x <this.MaxX && y>=0 && y <this.MaxY{
		Adjaoent = append(Adjaoent,&this.points[x][y])
	}
	if x,y := curPoint.x+1 , curPoint.y ;x>=0 && x <this.MaxX && y>=0 && y <this.MaxY{
		Adjaoent = append(Adjaoent,&this.points[x][y])
	}
	if x,y := curPoint.x+1 , curPoint.y+1 ;x>=0 && x <this.MaxX && y>=0 && y <this.MaxY{
		Adjaoent = append(Adjaoent,&this.points[x][y])
	}
	if x,y := curPoint.x , curPoint.y+1 ;x>=0 && x <this.MaxX && y>=0 && y <this.MaxY{
		Adjaoent = append(Adjaoent,&this.points[x][y])
	}
	if x,y := curPoint.x-1 , curPoint.y+1 ;x>=0 && x <this.MaxX && y>=0 && y <this.MaxY{
		Adjaoent = append(Adjaoent,&this.points[x][y])
	}
	if x,y := curPoint.x-1 , curPoint.y ;x>=0 && x <this.MaxX && y>=0 && y <this.MaxY{
		Adjaoent = append(Adjaoent,&this.points[x][y])
	}
	if x,y := curPoint.x-1 , curPoint.y-1 ;x>=0 && x <this.MaxX && y>=0 && y <this.MaxY{
		Adjaoent = append(Adjaoent,&this.points[x][y])
	}

	return Adjaoent
}

func (this *Map)printMap (path *SearchRoad){
	fmt.Println("地图的边际",this.MaxX,this.MaxY)
	for x := 0 ;x<this.MaxX;x++{
		for y := 0 ;y<this.MaxY;y++{
			if path!=nil{
				if x == path.start.x && y == path.start.y{
					fmt.Printf("%2s","S")
					goto NEXT
				}
				if x == path.end.x && y == path.end.y{
					fmt.Printf("%2s","E")
					goto NEXT
				}
				for i:= 0 ; i< len(path.TheRoad);i++{
					if path.TheRoad[i].x==x &&path.TheRoad[i].y==y{
						fmt.Printf("%2s","*")
						goto NEXT
					}
				}
			}
			fmt.Printf("%2s",this.points[x][y].view)

		NEXT:
		}
	}
}