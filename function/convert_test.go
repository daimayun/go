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

func TestFenToYuanToString(t *testing.T) {
	if yuan := FenToYuanToString(1); yuan != "0.01" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(301); yuan != "3.01" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(9); yuan != "0.09" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(10091); yuan != "100.91" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(12345678900); yuan != "123456789" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(163); yuan != "1.63" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(492); yuan != "4.92" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(9999999); yuan != "99999.99" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(9000099997899); yuan != "90000999978.99" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(100); yuan != "1" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(10); yuan != "0.1" {
		t.Error("FenToYuanToString() error.")
	}
	if yuan := FenToYuanToString(314); yuan != "3.14" {
		t.Error("FenToYuanToString() error.")
	}
}
