package ui

import (
	"fmt"
	"github.com/kalpio/audiodiscover/domain"
)

func ChooseDevice(devices []domain.Device) (*domain.Device, error) {
	fmt.Println("Choose audio device:")
	exitIndex := len(devices) + 1
	for i, d := range devices {
		fmt.Printf("\t%d. %s\n", i+1, d.Name)
	}
	fmt.Printf("\t%d. Exit\n", exitIndex)

	chosenValue := -1

	isValidValue := func() bool {
		return chosenValue > 0 && chosenValue <= exitIndex
	}

	for !isValidValue() {
		fmt.Printf("Enter device index (or type %d to exit): ", exitIndex)
		_, err := fmt.Scanln(&chosenValue)
		if err != nil {
			chosenValue = -1
		}
	}

	if chosenValue == exitIndex {
		return nil, nil
	}

	return &devices[chosenValue-1], nil
}
