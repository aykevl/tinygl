package style

type Scale uint8

func NewScale(percent int) Scale {
	return Scale(percent / 25)
}

