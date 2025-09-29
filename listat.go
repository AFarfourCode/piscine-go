package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	index := 0
	for l != nil {
		if index == pos {
			return l
		}
		l = l.Next
		index++
	}
	return nil
}
