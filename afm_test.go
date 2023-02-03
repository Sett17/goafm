package goafm

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.ReadFile("Times-Roman.afm")
	if err != nil {
		t.Fatal(err)
	}
	afm, err := Parse(file)
	if err != nil {
		t.Fatal(err)
	}
	if afm.FontName != "Times-Roman" {
		t.Errorf("FontName = %v, want %v", afm.FontName, "Times-Roman")
	}
	if afm.Weight != "Roman" {
		t.Errorf("Weight = %v, want %v", afm.Weight, "Roman")
	}
	if afm.ItalicAngle != 0 {
		t.Errorf("ItalicAngle = %v, want %v", afm.ItalicAngle, 0)
	}
	if afm.IsFixedPitch != false {
		t.Errorf("IsFixedPitch = %v, want %v", afm.IsFixedPitch, false)
	}
	if afm.UnderlinePosition != -100 {
		t.Errorf("UnderlinePosition = %v, want %v", afm.UnderlinePosition, -100)
	}

	if len(afm.CharMetricsByCode) == 0 {
		t.Errorf("CharMetricsByCode = %v, want %v", len(afm.CharMetricsByCode), "not empty")
	}
	if len(afm.CharMetricsByName) != 315 {
		t.Errorf("CharMetricsByName = %v, want %v", len(afm.CharMetricsByName), "315")
	}

	KernPairSize := 0
	for _, v := range afm.KernPair {
		KernPairSize += len(v)
	}
	if KernPairSize != 2073 {
		t.Errorf("KernPairSize = %v, want %v", KernPairSize, "2073")
	}
}

func BenchmarkParse(b *testing.B) {
	file, err := os.ReadFile("Times-Roman.afm")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Parse(file)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ReportAllocs()
}
