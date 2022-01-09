// 1041 Robot Bounded In Circle
//
// https://leetcode.com/problems/robot-bounded-in-circle/

// isRobotBounded function returns whether the robot's movement is in the boundary of circle or not.
// Robot can move straight forward or turn left or turn right with the input instructions.
// This is quite simple that robot can be in the circle boundary if robot is back to the origin point or looking at the different direction when the instructions are finished.
// On the former case, robot is always back to where its starting point.
// On the latter case, robot will be back to where its starting point by iterating the same instructions because it looks different direction.
// Thus, robot cannot be in the circle boundary if the robot looks the same direction with the initial time after the instructions.
// The robot's initial direction is index 0 on the code.
func isRobotBounded(instructions string) bool {
	displacement := []int{1, 1, -1, -1}
	x, y, direction := 0, 0, 0
	for _, i := range instructions {
		if i == 'L' {
			direction = (direction + 1) % 4
		} else if i == 'R' {
			direction = (direction + 3) % 4
		} else {
			if direction%2 == 0 {
				y += displacement[direction]
			} else {
				x += displacement[direction]
			}
		}
	}
	return x == 0 && y == 0 || direction != 0
}
