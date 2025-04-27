package controller

import (
	"fyne.io/fyne/v2"
	hook "github.com/robotn/gohook"
)

type ShortcutController struct {
	Window       fyne.Window
	WindowHidden bool
}

func NewShortcutController(w fyne.Window) *ShortcutController {
	return &ShortcutController{Window: w}
}

func (c *ShortcutController) GlobalShortcutsListener() {
	hook.Register(hook.KeyHold, []string{"ctrl", "alt", "h"}, func(e hook.Event) {
		fyne.DoAndWait(func() {
			if c.WindowHidden {
				c.Window.Show()
				c.Window.RequestFocus()
			} else {
				c.Window.Hide()
			}

			c.WindowHidden = !c.WindowHidden
		})
	})

	s := hook.Start()
	<-hook.Process(s)
}
