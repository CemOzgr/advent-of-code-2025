package day5

type Ranges []*Range

type Range struct {
	start int
	end   int
}

func (r *Range) ShouldUnite(other *Range) bool {
	var maxStart, minimumEnd = r.start, r.end

	if other.start > r.start {
		maxStart = other.start
	}

	if other.end < r.end {
		minimumEnd = other.end
	}

	return maxStart <= minimumEnd+1
}

func (r *Range) Unite(other *Range) {
	var newStart, newEnd = other.start, other.end
	if r.start <= other.start {
		newStart = r.start
	}

	if r.end >= other.end {
		newEnd = r.end
	}

	r.start = newStart
	r.end = newEnd
}

func (ranges *Ranges) Add(newRange *Range) {
	var united bool

	for _, r := range *ranges {
		if r.ShouldUnite(newRange) {
			r.Unite(newRange)
			united = true

			break
		}
	}

	if !united {
		*ranges = append(*ranges, newRange)
	}
}
