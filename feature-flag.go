if feature.IsEnabled("newWidget") {
	widget := widget.New()
} else {
	widget := widget.Old()
}

widget.render()
