// +build darwin

package desktop

type NSPoint CGPoint

func NSPointNew(x, y float64) NSPoint {
	p := NSPoint{CGFloat(x), CGFloat(y)}
	return p
}
