package utils

import "testing"

func Test(t *testing.T) {
	t.Run("test check dates", func(t *testing.T) {
		isValid := CheckDates(21, "jan")
		if !isValid {
			t.Errorf("should be valid")
		}
		isValid = CheckDates(31,"feb")
		if isValid {
			t.Errorf("should not be valid")
		}
	})
	t.Run("test check names", func(t *testing.T) {
		isValid := ValidName("test")
		if !isValid {
			t.Errorf("should be valid")
		}
		isValid = ValidName("test123")
		if isValid {
			t.Errorf("should not be valid")
		}
	})
	t.Run("test check mobile", func(t *testing.T) {
		isValid := ValidMobile(1234567890)
		if !isValid {
			t.Errorf("should be valid")
		}
		isValid = ValidMobile(1)
		if isValid {
			t.Errorf("should not be valid")
		}
	})
}
