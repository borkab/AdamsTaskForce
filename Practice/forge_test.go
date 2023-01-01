package ForgedInFire

import (
	"reflect"
	"testing"
)

func TestForge(t *testing.T) {
	damast := []string{"canister", "sanmai", "wave"}
	knife := []string{"dagger", "bowie", "recurve"}

	got := Forge(damast, knife)
	want := []string{"canister", "sanmai", "wave", "dagger", "bowie", "recurve"}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("the forge can't be with you")
	}
}
