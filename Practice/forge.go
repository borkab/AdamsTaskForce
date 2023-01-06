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

func Merge(sliceIN1, sliceIN2 []string) (sliceOUT []string) {
	m := make(map[string]bool)
	sliceIN1 = append(sliceIN1, sliceIN2...)

	for _, this := range sliceIN1 {
		if _, that := m[this]; !that {
			m[this] = true
			sliceOUT = append(sliceOUT, this)
		}
	}
	return sliceOUT
}

func Deduplicate(s []string) []string {
	m := make(map[string]bool)  //csinalj egy map-et string key es boolean ertek parokkal
	dedupSlice := []string{}    //a leendo listank duplikalt elemek nelkul
	for _, element := range s { //menj vegig az eredeti lista osszes elemen. amikor az elso elemhez ertel, nezd meg, h
		if _, v := m[element]; !v { //ez az elem szerepel-e mar a map kulcsai kozott
			m[element] = true                        //akkor tedd bele a mapba ezt az elemet, es
			dedupSlice = append(dedupSlice, element) //akkor add hozza a leendo listadhoz is ezt az elemet
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
