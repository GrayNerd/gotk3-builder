package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	ui "graynerd.com/gotk3-builder"
)

const bFile = "example.ui"

func main() {
	gtk.Init(nil)
	ui.LoadBuilder(bFile)
	win := ui.GetWindow("main_window")

	signals := map[string]interface{}{
		"destroy":                 func() { gtk.MainQuit() },
		"on_load_button_clicked":  func() { loadData() },
		"on_clear_button_clicked": func() { clearData() },
		"on_get_button_clicked":   func() { getData() },
		"on_close_button_clicked": func() { gtk.MainQuit() },
	}
	ui.ConnectSignals(signals)
	win.ShowAll()
	gtk.Main()
}

func loadData() {
	ui.GetEntry("entry_1").SetText("Entry One")
	ui.GetEntry("entry_2").SetText("Entry Two")
	ui.GetEntry("entry_3").SetText("Entry Three")
	ui.GetComboBox("combo_1").SetActive(4)
}

func clearData() {
	ui.GetEntry("entry_1").SetText("")
	ui.GetEntry("entry_2").SetText("")
	ui.GetEntry("entry_3").SetText("")
	ui.GetComboBox("combo_1").SetActive(-1)
}

func getData() {
	e1, err := ui.GetEntry("entry_1").GetText()
	if err != nil {
		log.Println(err.Error())
	}
	e2, err := ui.GetEntry("entry_2").GetText()
	if err != nil {
		log.Println(err.Error())
	}
	e3, err := ui.GetEntry("entry_3").GetText()
	if err != nil {
		log.Println(err.Error())
	}

	c1 := ui.GetComboBox("combo_1").GetActive()

	ce, err := ui.GetEntry("combo_entry").GetText()
	if err != nil {
		log.Println(err.Error())
	}

	dlg := gtk.MessageDialogNew(ui.GetWindow("main_window"), gtk.DIALOG_DESTROY_WITH_PARENT,
		gtk.MESSAGE_INFO, gtk.BUTTONS_OK,
		"entry_1: %s\nentry_2: %s\nentry_3: %s\ncombo value: %d, combo entry: %s", e1, e2, e3, c1, ce)

	x := dlg.Run()
	dlg.Destroy()
	log.Println(x)
}
