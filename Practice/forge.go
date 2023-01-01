package ForgedInFire

func Forge(a []string, b []string) []string {
	var c []string
	c = append(c, a...)
	c = append(c, b...)
	return c
}
