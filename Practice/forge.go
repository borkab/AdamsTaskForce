package ForgedInFire

func Forge(a []string, b []string) []string {
	var c []string
	c = append(c, a...)
	c = append(c, b...)
	return c
}

func Coppy(a []string) *[]string {
	return &a
}

func IfHasFire(a []string, b string) bool {
	var hasFire bool
	hasFire = false
	for _, v := range a {
		if v == b {
			hasFire = true
			break
		}
	}
	return hasFire
}
