package GolandProjects

import (
	potfile "Potfile/internal"
	"Potfile/internal/mask"
)

func GetRawString() string {
	return potfile.RawString()
}

func GetRawBytes() []byte {
	return potfile.RawBytes()
}

func GetParsedMap() []map[string]string {
	return potfile.ParsedMap()
}

func GetPasswords() []string {
	return potfile.Passwords()
}

func GetLocation() string {
	return potfile.Location()
}

func GetCount() int {
	return potfile.Count()
}

func GetMasks() []string {
	var masks []string
	passwords := potfile.Passwords()

	for _, password := range passwords {
		masks = append(masks, mask.GenerateMask(password))
	}
	return masks
}

func GetTopMasks(top int) []string {
	var masks []string

	passwords := potfile.Passwords()
	manager := mask.NewManager()
	for _, password := range passwords {
		manager.Add(password)
	}
	manager.Analyze()
	masks = manager.GetTopMasks(top)
	return masks
}

func WriteTopMasksToFile(top int, directory string) {
	passwords := potfile.Passwords()
	manager := mask.NewManager()
	for _, password := range passwords {
		manager.Add(password)
	}
	manager.Analyze()
	manager.WriteTopMasksToFile(top, directory)
}
