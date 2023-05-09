package function

import "testing"

func TestStringYuanToInt64Fen(t *testing.T) {
	if fen, err := StringYuanToInt64Fen("0.01"); err != nil || fen != 1 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("3.01"); err != nil || fen != 301 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("0.09"); err != nil || fen != 9 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("100.91"); err != nil || fen != 10091 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("123456789"); err != nil || fen != 12345678900 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("1.63"); err != nil || fen != 163 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("4.92"); err != nil || fen != 492 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("99999.99"); err != nil || fen != 9999999 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("90000999978.99"); err != nil || fen != 9000099997899 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("1.00"); err != nil || fen != 100 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("1.0"); err != nil || fen != 100 {
		t.Error("StringYuanToInt64Fen() error.")
	}
	if fen, err := StringYuanToInt64Fen("3.14"); err != nil || fen != 314 {
		t.Error("StringYuanToInt64Fen() error.")
	}
}

func TestYuanToInt64Fen(t *testing.T) {
	if fen, err := YuanToInt64Fen(0.01); err != nil || fen != 1 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(3.01); err != nil || fen != 301 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(0.09); err != nil || fen != 9 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(100.91); err != nil || fen != 10091 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(123456789); err != nil || fen != 12345678900 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(1.63); err != nil || fen != 163 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(4.92); err != nil || fen != 492 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(99999.99); err != nil || fen != 9999999 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(90000999978.99); err != nil || fen != 9000099997899 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(1.00); err != nil || fen != 100 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(1.0); err != nil || fen != 100 {
		t.Error("YuanToInt64Fen() error.")
	}
	if fen, err := YuanToInt64Fen(3.14); err != nil || fen != 314 {
		t.Error("YuanToInt64Fen() error.")
	}
}
