package main
 
func Intersection (p0, p1, p2, p3 PublicMatchState_Vector2Df) (bool, PublicMatchState_Vector2Df) {
	s1_x := p1.X - p0.X
	s1_y := p1.Y - p0.Y
	s2_x := p3.X - p2.X
	s2_y := p3.Y - p2.Y

	s := (-s1_y * (p0.X - p2.X) + s1_x * (p0.Y - p2.Y)) / (-s2_x * s1_y + s1_x * s2_y);
	t := ( s2_x * (p0.Y - p2.Y) - s2_y * (p0.X - p2.X)) / (-s2_x * s1_y + s1_x * s2_y);
	
	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
        return true, PublicMatchState_Vector2Df {X: p0.X + (t * s1_x), Y: p0.Y + (t * s1_y)};
    }

    return false, PublicMatchState_Vector2Df {}; 
}


func IntersectingBorders (start *PublicMatchState_Vector2Df, target *PublicMatchState_Vector2Df, m *Map) (bool) {
	for _, b := range m.Borders {
		hasIntersection, _ := Intersection(*start, *target, b.A, b.B)
		if hasIntersection {
			return true
		}
	}
	return false
}