package main

// 超时
//type WordDictionary struct {
//	wordsMap map[string]struct{}
//}
//
//func Constructor() WordDictionary {
//	return WordDictionary{
//		wordsMap: map[string]struct{}{},
//	}
//}
//
//func (w *WordDictionary) AddWord(word string) {
//	w.wordsMap[word] = struct{}{}
//}
//
//func (w *WordDictionary) Search(word string) bool {
//	for words := range w.wordsMap {
//		count := 0
//
//		if len(words) != len(word) {
//			continue
//		}
//
//		for i := 0; i < len(word); i++ {
//
//			if string(word[i]) == "." {
//				count++
//			} else if words[i] == word[i] {
//				count++
//			}
//		}
//
//		if count == len(words) {
//			return true
//		}
//	}
//	return false
//}

type TrieNode struct {
	children [26]*TrieNode
	IsEnd    bool
}

func (t *TrieNode) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
	}
	node.IsEnd = true
}

type WordDictionary struct {
	treeNode *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{treeNode: &TrieNode{}}
}

func (w *WordDictionary) AddWord(word string) {
	w.treeNode.Insert(word)
}

func (w *WordDictionary) Search(word string) bool {
	var dfs func(index int, node *TrieNode) bool
	dfs = func(index int, node *TrieNode) bool {
		if index == len(word) {
			return node.IsEnd
		}
		ch := word[index]
		if ch != '.' {
			child := node.children[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range node.children {
				child := node.children[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, w.treeNode)
}

func main() {

}
