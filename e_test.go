package sego

import (
	"testing"
)

func TestSego(t *testing.T) {
	dict := NewDictionary() //sego 存在bug，如果添加的token为文本的开始部分，则不进入分词

	dict.AddToken(*NewToken("美利坚", "$GW", 5))
	dict.AddToken(*NewToken("银行", "$ST", 100))
	dict.AddToken(*NewToken("人民", "$ST", 100))
	seg := NewSegmenter(dict)

	sgs := seg.Segment([]byte("美利坚人民银行"))

	for _, sg := range sgs {
		token := sg.Token()

		t.Log(token.Text(), ":", token.Pos())
	}
}

//func TestOSego(t *testing.T) {
//	var segmenter sego.Segmenter
//	segmenter.LoadDictionary("/data/dic/sego/t.dic")
//	// 分词
//	text := []byte("美利坚人民银行")
//	sgs := segmenter.Segment(text)

//	for _, sg := range sgs {
//		token := sg.Token()

//		t.Log(token.Text(), ":", token.Pos())
//	}
//}
