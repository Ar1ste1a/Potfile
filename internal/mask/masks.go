package mask

type Pattern int64

const (
	lowercase Pattern = iota
	uppercase
	digit
	symbol
	unknown
)

const (
	l = "abcdefghijklmnopqrstuvwxyz"
	u = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	d = "0123456789"
	h = "0123456789abcdef"
	H = "0123456789ABCDEF"
	s = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	a = "?l?u?d?s"
)

type Mask struct {
	mask string
	hits int
	len  int
}

func NewMask(password string) *Mask {
	return &Mask{mask: GenerateMask(password), len: len(password)}
}

func (m *Mask) Hit() {
	m.hits++
}

func (m *Mask) Len() int {
	return m.len
}

func (m *Mask) String() string {
	return m.mask
}

func GenerateMask(s string) string {
	mask := ""
	for _, r := range s {
		if r >= 48 && r < 58 {
			mask += "?d"
		} else if r > 64 && r < 91 {
			mask += "?u"
		} else if r > 96 && r < 123 {
			mask += "?l"
		} else {
			mask += "?s"
		}
	}
	return mask
}
