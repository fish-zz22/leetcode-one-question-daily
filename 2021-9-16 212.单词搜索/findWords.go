package main

import "fmt"

func main() {
	board := [][]byte{
		{'o','a','a','n'},
		{'e','t','a','e'},
		{'i','h','k','r'},
		{'i','f','l','v'},
	}
	words := []string{"oath","pea","eat","rain"}
	res := findWords(board,words)
	fmt.Println(res)
}

func findWords(board [][]byte, words []string) []string {
	m := len(board)
	n := len(board[0])

	t := &Trie{}
	// 初始化字典树
	for _,word := range words{
		t.Add(word)
	}


	res := make([]string,0)
	dirs := [][]int{{0,1},{0,-1},{1,0},{-1,0}}

	// 判断是否被重复使用
	used := make([][]bool,m)
	for i:=0;i<m;i++{
		used[i] = make([]bool,n)
	}

	var dfs func(x,y int, node *Trie)
	dfs = func(x, y int, node *Trie) {

		// 提前剪枝  1.当前node不匹配    2.第一个字母不匹配
		if node == nil || node.count == 0{
			return
		}

		// 判断是否来过
		if used[x][y] == true{
			return
		}
		used[x][y] = true

		if node.isEnd == true{
			res = append(res,node.word)
			// 匹配到就删除,防止有重复的字符并提高剪枝的效率
			t.Del(node.word)
		}

		for _,v := range dirs{
			i,j := x+v[0],y+v[1]
			if i>=0 && i<m && j>=0 && j<n {
				ch := board[i][j] - 'a'
				dfs(i,j,node.children[ch])
			}
		}
		used[x][y] = false
	}

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			dfs(i,j,t.children[board[i][j]-'a'])
		}
	}

	return res
}

type Trie struct {
	children [26]*Trie
	isEnd bool	// 判断是否到最后
	count int  //计数
	word string	//匹配的word 记录
}

func (t *Trie) Add(word string)  {
	node := t
	for _,ch := range word{
		ch = ch - 'a'
		if node.children[ch] == nil{
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
		node.count++
	}
	node.isEnd = true
	node.word =word
}

func (t *Trie) Del(word string)  {
	node := t
	for _,ch := range word{
		ch = ch - 'a'
		node = node.children[ch]
		node.count --
	}
	node.isEnd = false
}