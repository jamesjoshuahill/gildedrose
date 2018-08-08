package inventory

const (
	minQuality = 0
	maxQuality = 50
)

type quality struct {
	value int
}

func (q *quality) zero() {
	q.value = 0
}

func (q *quality) update(amount int) {
	v := q.value + amount

	if v < minQuality {
		v = minQuality
	}
	if v > maxQuality {
		v = maxQuality
	}

	q.value = v
}
