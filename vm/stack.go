package vm

type Stack struct {
	data []float64
}

func (s *Stack) Push(value float64) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() float64 {
	if len(s.data) == 0 {
		return 0.0 // 错误处理
	}
	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]
	return value
}
