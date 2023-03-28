package fs

import "testing"

// Not yet a proper test, just to check if this is working
func TestMaybeCreateProgramFolder(t *testing.T) {
	err := MaybeCreateProgramFolder("testingShit")
	if err != nil {
		t.Errorf("MyType.AppendToMessage did not update the message correctly. Error: %v", err)
	}
}
