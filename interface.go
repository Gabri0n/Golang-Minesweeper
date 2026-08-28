// User interface

package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func DrawBoard(table *tview.Table, board *Board) {
	for x, row := range board.Cells {
		for y, val := range row {
			color := tcell.ColorWhite
			BackgroundColour := tcell.ColorDarkGrey

			var bombColors = map[int]tcell.Color{
				1: tcell.ColorBlue,
				2: tcell.ColorGreen,
				3: tcell.ColorRed,
				4: tcell.ColorDarkBlue,
				5: tcell.ColorDarkRed,
				6: tcell.ColorTeal,
				7: tcell.ColorBlack,
				8: tcell.ColorGray,
			}

			var text string
			switch {
			case val.isFlagged: // Flag
				text = "?"
				color = tcell.ColorRed

			case val.isMine && !val.isRevealed: // Unrevealed Bomb
				text = "#"
				color = tcell.ColorWhite

			case val.isMine: // Revealed Bomb
				text = "*"
				color = tcell.ColorBlack
				BackgroundColour = tcell.ColorRed

			case val.isRevealed && val.adjacentBombs > 0: // Adjacent Bomb Count
				text = strconv.Itoa(val.adjacentBombs)
				color = bombColors[val.adjacentBombs]

			case val.isRevealed: // Revealed Empty Cell
				text = " "
				color = tcell.ColorBlack

			default: // Default Cell
				text = "#"
				color = tcell.ColorWhite

			}

			table.SetCell(x, y, tview.NewTableCell(text).SetTextColor(color).SetBackgroundColor(BackgroundColour).SetAlign(tview.AlignCenter).SetMaxWidth(1))

		}
	}
}

func RenderTui(board *Board) {

	tview.Borders.Horizontal = tview.BoxDrawingsLightHorizontal
	tview.Borders.Vertical = tview.BoxDrawingsLightVertical
	tview.Borders.TopLeft = tview.BoxDrawingsLightArcDownAndRight
	tview.Borders.TopRight = tview.BoxDrawingsLightArcDownAndLeft
	tview.Borders.BottomLeft = tview.BoxDrawingsLightArcUpAndRight
	tview.Borders.BottomRight = tview.BoxDrawingsLightArcUpAndLeft

	tview.Borders.HorizontalFocus = tview.BoxDrawingsLightHorizontal
	tview.Borders.VerticalFocus = tview.BoxDrawingsLightVertical
	tview.Borders.TopLeftFocus = tview.BoxDrawingsLightArcDownAndRight
	tview.Borders.TopRightFocus = tview.BoxDrawingsLightArcDownAndLeft
	tview.Borders.BottomLeftFocus = tview.BoxDrawingsLightArcUpAndRight
	tview.Borders.BottomRightFocus = tview.BoxDrawingsLightArcUpAndLeft

	app := tview.NewApplication()

	Table := tview.NewTable().SetBorders(false).SetSeparator(' ')

	Header := tview.NewTextView().SetText("SCORE HERE														TIMER HERE").SetTextColor(tcell.ColorWhite).SetTextAlign(tview.AlignCenter).SetScrollable(false)
	Header.ScrollToBeginning()
	Header.SetBorder(true).SetTitle(" Golang Minesweeper ")

	Field := tview.NewFlex().SetDirection(tview.FlexColumn)
	Field.SetBorder(true).SetBorderColor(tcell.ColorWhite)
	FieldWidth := 2 * len(board.Cells[0])

	spacer := tview.NewBox().SetBackgroundColor(tcell.ColorBlack)

	Field.AddItem(spacer, 0, 1, false).AddItem(Table, FieldWidth, 0, true).AddItem(spacer, 0, 1, false)

	boxContainer := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(Header, 3, 0, false).AddItem(Field, 0, 1, true)

	Table.SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorRed))
	Table.SetSelectable(true, true)

	DrawBoard(Table, board)
	Table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		x, y := Table.GetSelection()

		switch event.Rune() {

		case 'f': // flag
			board.Cells[x][y].ToggleFlagged()
			DrawBoard(Table, board)
			return nil

		case ' ':
			board.Select(x, y)
			DrawBoard(Table, board)
			return nil
		}

		if event.Key() == tcell.KeyEscape {
			app.Stop()
			return nil
		}

		return event
	})

	if err := app.SetRoot(boxContainer, true).SetFocus(Table).Run(); err != nil {
		panic(err)
	}
}
