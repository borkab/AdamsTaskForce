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

/*
func Merge(dst, src []string) (copper []string) {

	copy(dst, src) //masold src tartalmat dst-be

	for _, v := range dst { //menj vegig dst osszes elemen. amikor dst elso elemehez ertel,
		if !Contains(copper, v) { //nezd meg, h copper tartalmazza e mar dst elso elemet.
			copper = append(copper, v) //ha nem, akkor add hozza copperhez dst elso elemet
		}
	}
	return copper
}

*/

func Deduplicate(s []string) []string {
	m := make(map[string]bool)  //csinalj egy map-et string key es boolean ertek parokkal
	dedupSlice := []string{}    //a leendo listank duplikalt elemek nelkul
	for _, element := range s { //menj vegig az eredeti lista osszes elemen. amikor az elso elemhez ertel, nezd meg, h
		if _, v := m[element]; !v { //ha m
			m[element] = true                        //
			dedupSlice = append(dedupSlice, element) //akkor add hozze a leendo listadhoz ezt az elemet
		}
	}
	return dedupSlice
}

func Contains(slice []string, value string) bool {
	for _, s := range slice {
		if value == s {
			return true
		}
	}
	return false
}
