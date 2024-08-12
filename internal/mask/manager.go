package mask

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
)

type Manager struct {
	masks     map[string]*Mask
	lenMap    map[int][]Mask
	bestMasks map[int][]string
	total     int
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Add(password string) {
	if m.masks == nil {
		m.masks = make(map[string]*Mask)
	}

	mask := NewMask(password)

	if maskValue, exists := m.masks[mask.String()]; exists {
		maskValue.Hit()
		return
	} else {
		m.masks[mask.String()] = mask
	}
	m.total++
}

func (m *Manager) Total() int {
	return m.total
}

func (m *Manager) Analyze() {
	if m.lenMap == nil {
		m.lenMap = make(map[int][]Mask)
	}

	for _, mask := range m.masks {
		m.lenMap[mask.Len()] = append(m.lenMap[mask.Len()], *mask)
	}

	for l, masks := range m.lenMap {
		if l > 0 {
			m.writeMasksToFile(l, masks)
		}
	}
}

func (m *Manager) PrintMasks() {
	for _, mask := range m.masks {
		println(mask.String())
	}
}

func (m *Manager) writeMasksToFile(l int, masks []Mask) {
	filename := fmt.Sprintf("%d-char-mask", l)
	file, err := os.Create(filename)
	if err != nil {
		println("Error: Cannot create mask file")
		os.Exit(1)
	}

	// Sort masks by hits in descending order
	sort.Slice(masks, func(i, j int) bool {
		return masks[i].hits > masks[j].hits
	})

	defer file.Close()

	for _, mask := range masks {
		line := fmt.Sprintf("%s\n", strings.TrimSpace(mask.String()))
		file.WriteString(line)
	}
}

func (m *Manager) writeTop10MasksToFile() {
	masks := make([]Mask, 0, len(m.masks))
	for _, mask := range m.masks {
		masks = append(masks, *mask)
	}

	file, err := os.Create("top-10-masks")
	if err != nil {
		println("Error: Cannot create mask file")
		os.Exit(1)
	}

	defer file.Close()

	// Sort masks by hits in descending order
	sort.Slice(masks, func(i, j int) bool {
		return masks[i].hits > masks[j].hits
	})

	for i, mask := range masks {
		if i == 10 {
			break
		}
		line := fmt.Sprintf("%s\n", mask.String())
		fmt.Printf("\n%d) %s - %d", i+1, mask.String(), mask.hits)
		file.WriteString(line)
	}
	fmt.Println()
}

func (m *Manager) WriteTopMasksToFile(top int, directory string) {
	masks := make([]Mask, 0, len(m.masks))
	for _, mask := range m.masks {
		masks = append(masks, *mask)
	}

	filename := fmt.Sprintf("top-%d-masks", top)
	filename = path.Join(directory, filename)
	file, err := os.Create(filename)
	if err != nil {
		println("Error: Cannot create mask file")
		os.Exit(1)
	}

	defer file.Close()

	// Sort masks by hits in descending order
	sort.Slice(masks, func(i, j int) bool {
		return masks[i].hits > masks[j].hits
	})

	for i, mask := range masks {
		if i == top {
			break
		}
		line := fmt.Sprintf("%s\n", mask.String())
		file.WriteString(line)
	}
}

func (m *Manager) GetTopMasks(top int) []string {
	out := make([]string, 0, top)

	masks := make([]Mask, 0, len(m.masks))
	for _, mask := range m.masks {
		masks = append(masks, *mask)
	}

	// Sort masks by hits in descending order
	sort.Slice(masks, func(i, j int) bool {
		return masks[i].hits > masks[j].hits
	})

	for i, mask := range masks {
		if i == top {
			break
		}
		out = append(out, mask.String())
	}
	return out
}

func (m *Manager) printTop25Masks() {
	fmt.Println()
	masks := make([]Mask, 0, len(m.masks))
	for _, mask := range m.masks {
		masks = append(masks, *mask)
	}

	// Sort masks by hits in descending order
	sort.Slice(masks, func(i, j int) bool {
		return masks[i].hits > masks[j].hits
	})

	for i, mask := range masks {
		if i == 25 {
			break
		}
		fmt.Printf("%d) %s - %d\n", i+1, mask.String(), mask.hits)
	}
}
