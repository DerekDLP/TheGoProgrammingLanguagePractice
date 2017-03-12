package main

import (
	"bytes"
	"fmt"
)

// IntSet :是一组小的非负整数
// 零值代表空集合
type IntSet struct {
	words []uint64
}

// Has 报告集合是否包含非负值x
func (s *IntSet) Has(x int) bool {
	// 因为每一个字含有64bit
	// x/64 作为字的下标
	// x%64 作为字内的bit的所在位置
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add 将非负值x添加到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith 设置S和T的并集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			// 一次性完成64个元素的'或'计算
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String 返回一个字符串数组，格式为"{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))

	// string方法使用的是指针接收器，所以下面的两个语句，只有第二个是正常输出的
	fmt.Println(x)
	fmt.Println(&x)
}
