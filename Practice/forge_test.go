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

func TestCoppy(t *testing.T) {
	knifes := []string{"canister", "sanmai", "wave", "dagger", "bowie", "recurve"}
	got := Coppy(knifes)
	want := &knifes

	if !reflect.DeepEqual(got, want) {
		t.Fatal("no coppy")
	}

}

func TestIfHasFire(t *testing.T) {
	knifes := []string{"canister", "sanmai", "wave", "dagger", "bowie", "recurve"}
	fire := "wave"
	got := IfHasFire(knifes, fire)
	want := true

	if got != want {
		t.Fatal("it has no fire")
	}
}
