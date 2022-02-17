package square

import (
	"testing"
)

func TestCalcSquare(t *testing.T) {
	var expected = 43.30127018922193
	if CalcSquare() != expected {
		t.Errorf("Ошибка. Expect %v, got %v",
			expected, CalcSquare())
	}
}
