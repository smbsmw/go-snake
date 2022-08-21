package snake

type snake struct {
	direction coord
	body      []coord
}

func (s *snake) head() coord {
	return s.body[len(s.body)-1]
}

func (s *snake) length() int {
	return len(s.body)
}

func (s *snake) move(body []coord) {
	h := s.head()
	s.body = append(body, coord{
		h.X + s.direction.X,
		h.Y + s.direction.Y,
	})
}

func (s *snake) nextMove() {
	s.move(s.body[1:])
}

func (s *snake) headOnBody() bool {
	for _, b := range s.body[:len(s.body)-1] {
		if s.head() == b {
			return true
		}
	}
	return false
}

func (s *snake) eat() {
	s.move(s.body)
}
