package main

import "fmt"

// import "strconv"

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) put(x, y int, u string) { // interface
	if u == "o" {
		b.tokens[x+3*y] = 1
	} else if u == "x" {
		b.tokens[x+3*y] = 2
	}
}

func (b *Board) get(x, y int) string { // interface
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else if b.tokens[x+3*y] == 2 {
		return "x"
	} else {
		return "-"
	}
}

//表示する関数
func (b *Board) display() {
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			//fmt.Println(b.tokens[i+3*j])
			fmt.Print(b.get(j, i))
		}
		fmt.Println("")
	}
}

//勝ち負けを判断する関数
// ○が３つ揃ってたら1を返す．xが３つ揃ってたら2を返す．何もなかったら0を返す.全て埋まっていたら3を返す．
func (b *Board) judge() int {
	var i, j int
	var tmp1, tmp2, tmp3 int
	var fullCheck int

	fullCheck = 0

	for i = 0; i < 9; i++ {
		if b.tokens[i] == 0 {
			fullCheck = 1
		}
	}

	if fullCheck == 0 {
		return 3
	}

	//縦に調べる　３通り
	for j = 0; j < 3; j++ {
		for i = j; i < 9; i += 3 {
			if i == j {
				tmp1 = b.tokens[i]
			} else if i == j+3 {
				tmp2 = b.tokens[i]
			} else if i == j+6 {
				tmp3 = b.tokens[i]
			}
		}

		if tmp1 == tmp2 && tmp2 == tmp3 {
			if tmp1 == 1 {
				return 1
			} else if tmp1 == 2 {
				return 2
			}
		}
	}

	//横に見よう！

	for j = 0; j < 9; j += 3 {
		for i = 0; i < 3; i++ {
			if i == 0 {
				tmp1 = b.tokens[j+0]
			} else if i == 1 {
				tmp2 = b.tokens[j+1]
			} else if i == 2 {
				tmp3 = b.tokens[j+2]
			}
		}

		if tmp1 == tmp2 && tmp2 == tmp3 {
			if tmp1 == 1 {
				return 1
			} else if tmp1 == 2 {
				return 2
			}
		}
	}

	//斜めに見よう

	for i = 0; i < 9; i += 4 {
		if i == 0 {
			tmp1 = b.tokens[i]
		} else if i == 4 {
			tmp2 = b.tokens[i]
		} else if i == 8 {
			tmp3 = b.tokens[i]
		}
	}

	if tmp1 == tmp2 && tmp2 == tmp3 {
		if tmp1 == 1 {
			return 1
		} else if tmp1 == 2 {
			return 2
		}
	}

	for i = 2; i < 7; i += 2 {
		if i == 2 {
			tmp1 = b.tokens[i]
		} else if i == 4 {
			tmp2 = b.tokens[i]
		} else if i == 6 {
			tmp3 = b.tokens[i]
		}
	}

	if tmp1 == tmp2 && tmp2 == tmp3 {
		if tmp1 == 1 {
			return 1
		} else if tmp1 == 2 {
			return 2
		}
	}

	return 0
}

func main() {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	var x, y int
	var c int
	c = 1

	//fmt.Print("Player 1: Input (x,y)")
	for {
		fmt.Print("Player")
		fmt.Print(c)
		fmt.Print(": Input (x,y) ")

		fmt.Scanf("%d %d", &x, &y)
		//Playerターン交代
		// c = 1：Player1, c = 2：Player2
		if b.get(y, x) == "-" {
			if c == 1 {
				b.put(y, x, "o")
				c = 2
			} else if c == 2 {
				b.put(y, x, "x")
				c = 1
			}
		} else {
			fmt.Println("既に置かれています!")
			b.display()
			continue
		}
		if b.judge() == 1 {
			fmt.Println("oが揃いました!")
			break
		}
		if b.judge() == 2 {
			fmt.Println("xが揃いました!")
			break
		}
		if b.judge() == 3 {
			fmt.Println("引き分けです！")
			break
		}

		b.display()
	}
	b.display()
}
