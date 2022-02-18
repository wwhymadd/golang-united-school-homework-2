package square

import (
	"testing"
)

func TestCalcSquare1(t *testing.T) {
	expected := 100.00
	if CalcSquare1(Square, SidesSquare) != expected {
		t.Errorf("Ошибка. Expect %f, got %f",
			expected, CalcSquare1(Square, SidesSquare))
	}
}

func TestCalcSquare2(t *testing.T) {
	expected := 43.250000
	if CalcSquare2(Triangle, SidesTriangle) != expected {
		t.Errorf("Ошибка. Expect %f, got %f",
			expected, CalcSquare2(Triangle, SidesTriangle))
	}
}

func TestCalcSquare3(t *testing.T) {
	expected := 314.150000
	if CalcSquare3(Circle, SidesCircle) != expected {
		t.Errorf("Ошибка. Expect %f, got %f",
			expected, CalcSquare3(Circle, SidesCircle))
	}
}
