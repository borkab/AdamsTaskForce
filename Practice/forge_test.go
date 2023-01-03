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

func TestCopy(t *testing.T) {
	src := []string{"canister", "sanmai", "wave", "dagger", "sanmai", "bowie", "recurve", "wave"}
	var dst []string

	got := Copy(src, dst)
	want := []string{"canister", "sanmai", "wave", "dagger", "bowie", "recurve"}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("I couldn't forge as you wish")
	}
}

func TestContains(t *testing.T) {
	slice := []string{"canister", "sanmai", "wave", "dagger", "bowie", "recurve"}
	str := "dagger"

	got := Contains(slice, str)
	want := true

	if got != want {
		t.Fatalf("got %#v want %#v", got, want)
	}
}
