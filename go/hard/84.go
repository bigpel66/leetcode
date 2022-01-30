// Largest Rectangle in Histogram
//
// https://leetcode.com/problems/largest-rectangle-in-histogram/

// CustomStack struct is used to manage the histogram to get the maximum area.
type CustomStack struct {
	Index    int   // current index of array heights
	MaxArea  int   // current calculater maximum area
	ArrValue []int // array of values in stack
	ArrIndex []int // array of indexes in stack
}

// largestRectangleArea function returns the maximum area of the given histogram.
// Idea is like these following 3 cases.
// If current bar is taller than previous one, push to the stack.
// Else if current bar is as same as previous one, nothing happens.
// Else if current bar is smaller than previous one, pop out from the stack.
// When the pop occurs, calculate the maximum area.
// Pop should not be the single case, but should be iterated until the current one is bigger than other previous ones.
// Popped elements will update the maximum area if it is ok.
func largestRectangleArea(heights []int) int {
	stack := &CustomStack{0, 0, make([]int, 0), make([]int, 0)}
	for _, currentHeight := range heights {
		stack.Index++
		previousHeieght, _ := stack.TopValueAndIndex()
		if previousHeieght < currentHeight {
			stack.PushValueAndIndex(currentHeight, stack.Index)
		}
		if previousHeieght > currentHeight {
			stack.Update(currentHeight)
		}
	}
	stack.Index++
	stack.Update(0)
	return stack.MaxArea
}

// Size function returns the length of the stack.
func (this *CustomStack) Size() int {
	return len(this.ArrValue)
}

// PushValueAndIndex function pushes the value and index to the stack.
func (this *CustomStack) PushValueAndIndex(value int, index int) {
	this.ArrValue = append(this.ArrValue, value)
	this.ArrIndex = append(this.ArrIndex, index)
}

// PopValueAndIndex function pops the value and index from the stack.
func (this *CustomStack) PopValueAndIndex() {
	if len(this.ArrValue) > 0 {
		this.ArrValue = this.ArrValue[:len(this.ArrValue)-1]
		this.ArrIndex = this.ArrIndex[:len(this.ArrIndex)-1]
	}
}

// TopValueAndIndex function returns the value and index of the top on the stack.
func (this *CustomStack) TopValueAndIndex() (int, int) {
	if this.Size() == 0 {
		return 0, 0
	}
	return this.ArrValue[len(this.ArrValue)-1], this.ArrIndex[len(this.ArrIndex)-1]
}

// CalculateArea function calculates the area of fixed point on the stack to the current top point on the stack.
func (this *CustomStack) CalculateArea() int {
	if this.Size() == 0 {
		return 0
	}
	value, index := this.TopValueAndIndex()
	return (this.Index - index) * value
}

// Update function updates the maximum area by popping the top value and index.
// When the calculating is done, pushing the current height will be processed.
func (this *CustomStack) Update(currentHeight int) {
	index := 0
	isFixedFirst := false
	for i := len(this.ArrValue) - 1; i >= 0; i-- {
		if this.ArrValue[i] > currentHeight {
			if !isFixedFirst {
				isFixedFirst = true
				if this.MaxArea < this.ArrValue[i] {
					this.MaxArea = this.ArrValue[i]
				}
			}
			index = this.ArrIndex[i]
			area := this.CalculateArea()
			this.PopValueAndIndex()
			if this.MaxArea < area {
				this.MaxArea = area
			}
		} else {
			break
		}
	}
	this.PushValueAndIndex(currentHeight, index)
}
