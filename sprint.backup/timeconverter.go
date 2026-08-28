package sprint

func TimeConverter(n int) (int, int, int) {
	h := int(n / 3600)
	m := int(((n - (h * 3600))) / 60)
	s := int(n - (h*3600) - (m*60))
	return h, m, s

}