package sego

//NewSegment 构建一个分词片结果
func NewSegment(start, end int, token Token) *Segment {
	return &Segment{start, end, &token}
}
