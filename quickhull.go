package quickhull

type Point struct {
	X, Y int
}

func Find(S []Point) (H []Point) {
	A, B := S[0], S[0]
	for _, P := range S[1:] {
		if P.X < A.X {
			A = P
		}
		if P.X > B.X {
			B = P
		}
	}
	var S1, S2 []Point
	dx, dy := B.X-A.X, B.Y-A.Y
	for _, P := range S {
		if A == P || B == P {
			continue
		}
		if dx*(P.Y-A.Y) < dy*(P.X-A.X) {
			S1 = append(S1, P)
		} else {
			S2 = append(S2, P)
		}
	}
	return append(append(append(append(H, A), findHull(S1, A, B)...), B), findHull(S2, B, A)...)
}

func findHull(Sk []Point, P, Q Point) (H []Point) {
	if len(Sk) == 0 {
		return nil
	}
	dx, dy := Q.X-P.X, Q.Y-P.Y
	C := Sk[0]
	d := dy*C.X - dx*C.Y
	for _, A := range Sk[1:] {
		if t := dy*A.X - dx*A.Y; t > d {
			C, d = A, t
		}
	}
	var S1, S2 []Point
	for _, A := range Sk {
		if A == C {
			continue
		}
		if (C.X-P.X)*(A.Y-C.Y) < (C.Y-P.Y)*(A.X-C.X) {
			S1 = append(S1, A)
		} else if (Q.X-C.X)*(A.Y-Q.Y) < (Q.Y-C.Y)*(A.X-Q.X) {
			S2 = append(S2, A)
		}
	}
	return append(append(append(H, findHull(S1, P, C)...), C), findHull(S2, C, Q)...)
}
