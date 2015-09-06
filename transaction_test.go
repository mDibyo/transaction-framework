package transaction

import "testing"

func TestNewAmountFigure(t *testing.T) {
	tests := []struct {
		units      int64
		fractions  int64
		debit      bool
		wantAmount float64
		wantError  bool
	}{
		{
			units:     -12,
			wantError: true,
		},
		{
			units:     12,
			fractions: -45,
			wantError: true,
		},
		{
			units:     12,
			fractions: 145,
			wantError: true,
		},
		{
			units:      12,
			fractions:  45,
			debit:      true,
			wantAmount: 12.45,
		},
		{
			units:      12,
			fractions:  45,
			wantAmount: -12.45,
		},
	}
	for _, test := range tests {
		af, gotErr := NewAmountFigure(test.units, test.fractions, test.debit)
		if test.wantError {
			if gotErr == nil {
				t.Errorf("did not get expected error.")
			}
		} else {
			if gotErr != nil {
				t.Errorf("got unexpected error: %v", gotErr)
			}
		}
		if gotAmount := af.Float64(); gotAmount != test.wantAmount {
			t.Errorf("did not get expected amount. actual=%v, expected=%v", gotAmount, test.wantAmount)
		}
	}
}

func TestRecordValid(t *testing.T) {
	tests := []struct {
		r         Record
		wantValid bool
	}{
		{
			r: Record{
				Debit{Amount: AmountFigure(1300)},
				Debit{Amount: AmountFigure(-1200)},
			},
			wantValid: false,
		},
		{
			r: Record{
				Debit{Amount: AmountFigure(1300)},
				Debit{Amount: AmountFigure(-1200)},
				Debit{Amount: AmountFigure(-50)},
				Debit{Amount: AmountFigure(-50)},
			},
			wantValid: true,
		},
	}
	for _, test := range tests {
		if gotValid := test.r.Valid(); gotValid != test.wantValid {
			t.Errorf("did not get expected validity. actual=%t, expected=%t", gotValid, test.wantValid)
		}
	}
}
