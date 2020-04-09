package main

import "fmt"

func main()  {
	prestMap := []string{
		". . . . . . . . . . . . . . . . . . \n",
		". . . . . . . . . . . . . . . . . . \n",
		". . . . . . . . . . . . . . . . . . \n",
		". . . . . . . . . . . . . . . . . . \n",
		". . . . . . . . . . . . . . . . . . \n",
		"X X . X X X X X X X X X X X X X X X \n",
		". . . . X . . . . . . . . . . . . . \n",
		". . . . X . X . . . . . . . . . . . \n",
		". . . . . . X . . . . . . . . . . . \n",
		"X X X X X X X . X X X X X X X X X X \n",
		". . . . . . . . . . . . . . . . . . \n",
		". . . . . . . . . . . . . . . . . . \n",
		"X X X X X X X X X . X X X X X X X X \n",
		". . . . . . . . . . . . . . . . . . \n",
		". . . . . . . . . . . . . . . . . . \n",
		". . . . . . . . . . . . . . . . . . \n",
		". . . . . . . . . . . . . . . . . . \n",
		". . . . . . . . . . . . . . . . . . \n",
		". . . . . . . . . . . . . . . . . . \n",
	}

	m := NewMap(prestMap)
	m.printMap(nil)
	searchroad := NewSearchRoad(0,0,17,17,&m)
	if searchroad.FindoutShortestPath(){
		fmt.Println("find it")
		m.printMap(searchroad)
	}else {
		fmt.Println("could not find it")
	}
}
