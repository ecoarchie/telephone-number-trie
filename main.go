package main

import (
	"fmt"
	"strconv"
)

type TrieNode struct {
		value rune
    children map[rune]*TrieNode
    isEndOfWord bool
}

func NewNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEndOfWord: false,
	}
}

type Trie struct {
    root *TrieNode
}


func Constructor() Trie {
    return Trie{
			root: NewNode(),
    }
}

func (this *TrieNode) AddChildNode(char rune) *TrieNode {
	_, ok := this.children[char]
	if !ok {
		this.children[char] = NewNode()
		this.children[char].isEndOfWord = true
		this.children[char].value = char
	}
	return this.children[char]
}


func (this *Trie) Insert(word string)  {
	cur := this.root
	for _, char := range word {
		_, ok := cur.children[char]
		if !ok {
			cur.children[char] = NewNode()
		}
		cur = cur.children[char]
	}
	cur.isEndOfWord = true
}


func (this *Trie) Search(word string) bool {
  cur := this.root
	for _, char := range word {
		_, ok := cur.children[char]
		if !ok {
			return false
		}
		cur = cur.children[char]
	} 
	return cur.isEndOfWord
}


func (this *Trie) StartsWith(prefix string) bool {
  cur := this.root
	for _, char := range prefix {
		_, ok := cur.children[char]
		if !ok {
			return false
		}
		cur = cur.children[char]
	} 
	return true
}

type Keyboard map[rune][]rune

func (kb Keyboard) Add(num rune, chars []rune) {
	kb[num] = chars
}

func NewKeyboard() *Keyboard {
	keyboard := make(Keyboard)
	keyboard.Add('2', []rune{'a', 'b', 'c'})
	keyboard.Add('3', []rune{'d', 'e', 'f'})
	keyboard.Add('4', []rune{'g', 'h', 'i'})
	keyboard.Add('5', []rune{'j', 'k', 'l'})
	keyboard.Add('6', []rune{'m', 'n', 'o'})
	keyboard.Add('7', []rune{'p', 'q', 'r', 's'})
	keyboard.Add('8', []rune{'t', 'u', 'v'})
	keyboard.Add('9', []rune{'w', 'x', 'y', 'z'})
	return &keyboard
}

func TrieFromTelNumber(keyboard Keyboard, node *TrieNode, telNum string) {
	if telNum == "" {
		return
	}
	for i, c := range telNum {
		for _, char := range keyboard[c] {
			childNode := node.AddChildNode(char)
			TrieFromTelNumber(keyboard, childNode, telNum[i+1:])
		}
	}

}

func Check(num int, words []string) []string {
	result := []string{}
	keyboard := NewKeyboard()
	telStr := strconv.Itoa(num)
	trie := Constructor()
	TrieFromTelNumber(*keyboard, trie.root, telStr)
	for _, w := range words {
		if trie.Search(w) {
			result = append(result, w)
		}
	}
	return result
}

func main() {
	res := Check(997336663, []string{"foo", "fooo", "bar", "baz", "foobar", "redmond"})
	// res := Check(27887, []string{"artur"})
	fmt.Println(res)

}