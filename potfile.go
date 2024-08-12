package GolandProjects

import (
	potfile "github.com/Ar1ste1a/Potfile/internal"
	"github.com/Ar1ste1a/Potfile/internal/mask"
)

func GetRawString() string {
	return potfile.RawString()
}

func GetRawBytes() []byte {
	return potfile.RawBytes()
}

func GetParsedMap() []map[string]string {
	return potfile.Map()
}

func GetPasswords() []string {
	return potfile.Passwords()
}

func GetHashes() []string {
	return potfile.Hashes()
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

func GetTopMasksByLength(length, top int) []string {
	var masks []string

	passwords := potfile.Passwords()
	manager := mask.NewManager()
	for _, password := range passwords {
		manager.Add(password)
	}
	manager.Analyze()
	masks = manager.GetTopMasksByLength(length, top)
	return masks
}

func WriteTopMasksByLengthToFile(length, top int, directory string) {
	passwords := potfile.Passwords()
	manager := mask.NewManager()
	for _, password := range passwords {
		manager.Add(password)
	}
	manager.Analyze()
	manager.WriteTopMasksToFileByLength(length, top, directory)
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

func SetPotfileLocation(location string) {
	potfile.SetLocation(location)
}
