package main

type Stack struct {
	Data []string
}

func (s *Stack) Push(x string) {
	s.Data = append(s.Data, x)
}
func (s *Stack) Pop() string {

	if l := len(s.Data); l > 0 {
		t := s.Data[l-1]
		s.Data = s.Data[:l-1]
		return t
	}
	panic("Can not pop anymore")
}

func main() {

}
