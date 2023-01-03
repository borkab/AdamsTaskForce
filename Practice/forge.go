/*
csinálj egy funkciót ami képes két slice -t egyesíteni

aztán csinálj egy funkciót ami képes egy slice alapján egy olyan másolatot készíteni,
 ami nem tartalmaz duplikált elemeket , viszont a sorrend alapvetően az elemek közt ,
 ugyan az mint a kiindulási slice ban volt

csinálj egy funkciót ami képes megmondani, ha egy string már benne egy string sliceban
*/

package ForgedInFire

func Forge(a []string, b []string) []string {
	var c []string

	c = append(c, a...)
	c = append(c, b...)
	return c
}

func Copy(src, dst []string) (copper []string) {

	copy(src, dst)

	for _, v := range dst {
		if !Contains(copper, v) {
			copper = append(copper, v)
		}
	}
	return copper
}

func Contains(slice []string, value string) bool {
	for _, s := range slice {
		if value == s {
			return true
		}
	}
	return false
}
