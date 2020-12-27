package game

import (
	"strconv"
)

//2048游戏
//规则，1.在n*n的格子中随机生成两个数字（2）哦

type Game struct {
	data []int
	list []string
	// indexstrconv.ltoa(i)
}

func (g *Game) self() {
	// g.index = make(map[string]int)
	var n = 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			g.list[n] = strconv.Itoa(i) + strconv.Itoa(j)
			n++
		}
	}
}

func (g *Game) staart() {

}
