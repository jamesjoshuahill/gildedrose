package inventory

const (
	minQuality = 0
	maxQuality = 50
)

type Quality struct {
	value int
}

func NewQuality(value int) *Quality {
	return &Quality{value: value}
}

func (q Quality) Value() int {
	return q.value
}

func (q *Quality) Update(amount int) {
	v := q.value + amount

	if v < minQuality {
		v = minQuality
	}
	if v > maxQuality {
		v = maxQuality
	}

	q.value = v
}
