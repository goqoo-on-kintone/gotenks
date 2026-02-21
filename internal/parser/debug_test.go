package parser

import (
	"fmt"
	"testing"
)

func TestDebugFacilityFile(t *testing.T) {
	result, err := ParseFile("../../testdata/facility-fields.d.ts")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	for _, iface := range result.Interfaces {
		fmt.Printf("Interface: %s (extends: %s)\n", iface.Name, iface.Extends)
		for i, f := range iface.Fields {
			fmt.Printf("  [%d] %s: type=%s, isSubtable=%v, subfields=%d\n",
				i, f.Name, f.TypeName, f.IsSubtable, len(f.SubtableFields))
		}
	}
}
