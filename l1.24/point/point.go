package point

import "math"

type Point struct {
	x, y float64
}

// Конструктор
func NewPoint(x, y float64) *Point { 
	return &Point{x, y}
}

// Метод для получения координаты X
func (p *Point) GetX() float64 {
	return p.x
}

// Метод для получения координаты Y
func (p *Point) GetY() float64 {
	return p.y
}

// Метод для установки координаты X
func (p *Point) SetX(x float64) {
	p.x = x
}

// Метод для установки координаты Y
func (p *Point) SetY(y float64) {
	p.y = y
}

// Расчет расстояния между точками
func (p1 *Point) GetDistance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x - p1.x, 2) + math.Pow(p2.y - p1.y, 2))
}