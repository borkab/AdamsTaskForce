package ForgedInFire

import (
	"reflect"
	"testing"

	"github.com/adamluzsi/testcase/assert"
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

/*
func TestMerge(t *testing.T) {
	src := []string{"canister", "sanmai", "wave", "dagger", "sanmai", "bowie", "recurve", "wave"}
	var dst []string

	got := Merge(dst, src)

	want := []string{"canister", "sanmai", "wave", "dagger", "bowie", "recurve"}

	assert.Equal(t, got, want)

	//	if !reflect.DeepEqual(got, want) {
	//		t.Fatal("I couldn't forge as you wish")
	//	}

}
*/

func TestDeduplicate(t *testing.T) {
	input := []string{"1", "3", "2", "4", "2", "3", "1", "5"}
	got := Deduplicate(input)

	assert.Equal(t, []string{"1", "3", "2", "4", "5"}, got)
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
