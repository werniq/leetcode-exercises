package main

import (
	"fmt"
	"strings"
)

// TrieNode represents is a node in Trie DS
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	words    []string
}

// T9Dictionary represents the T9 dictionary 
type T9Dictionary struct {
	root *TrieNode
}

// NewTrieNode creates a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
		words:    []string{},
	}
}

// NewT9Dictionary creates a new T9Dictionary
func NewT9Dictionary() *T9Dictionary {
	return &T9Dictionary{
		root: NewTrieNode(),
	}
}

// Insert adds a word to the Trie
func (t *T9Dictionary) Insert(word string) {
	node := t.root
	for _, char := range word {
		if child, ok := node.children[char]; ok {
			node = child
		} else {
			newNode := NewTrieNode()
			node.children[char] = newNode
			node = newNode
		}
	}
	node.isEnd = true
	node.words = append(node.words, word)
}

// Search returns words that match a given T9 sequence
func (t *T9Dictionary) Search(sequence string) []string {
	node := t.root
	for _, digit := range sequence {
		if child, ok := node.children[digit]; ok {
			node = child
		} else {
			return nil
		}
	}
	return node.words
}
