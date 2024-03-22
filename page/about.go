package page

import (
	"fmt"
	"schoperation/lethalloader/domain/viewer"
)

type AboutPage struct {
}

func NewAboutPage() AboutPage {
	return AboutPage{}
}

func (page AboutPage) Show(args any) (viewer.OptionsResult, error) {
	clear()

	fmt.Print(`
LethalLoader
Copyright (C) 2024 Schoperation

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

Lethal Company Copyright (C) 2023-24 Zeekerss
Thunderstore Copyright (C) 2024 Thunderstore Team

LethalLoader is in no way affiliated with Lethal Company or Thunderstore.`)

	fmt.Print("\n\n")
	fmt.Print("What to Do?\n")
	fmt.Print("-----------\n")
	fmt.Print("Q) Back to Main Menu\n")
	fmt.Print("\n")

	back := viewer.NewOption(viewer.OptionDto{
		Letter: 'Q',
		Page:   viewer.PageMainMenu,
	}, []string{})

	options := viewer.NewOptions([]viewer.Option{back})
	return options.TakeInput(), nil
}
