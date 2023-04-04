package style

type Scale uint8

func NewScale(percent int) Scale {
	return Scale((percent + 12) / 25)
}

func (s Scale) Percent() int {
	return int(s) * 25
}
