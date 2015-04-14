if user.IsInGroup("my-team") {
  widget := widget.New()
} else {
  widget := widget.Old()
}

widget.render();
