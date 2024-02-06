package dialog

import (
	"assistants/component"
	"asssistant/protocol/internal/protocol"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-faker/faker/v4"
	"github.com/samber/lo"
)

func requiresForm() (*fyne.Container, ConnectionDialogResult, *ConnectionForm) {
	result := ConnectionDialogResult{
		Protocol: binding.NewString(),
		Address:  binding.NewString(),
		Port:     binding.NewString(),
		Name:     binding.NewString(),
	}
	lo.Must0(result.Name.Set(faker.Word()))
	nameLabel := widget.NewLabel("Name")
	nameEntry := widget.NewEntryWithData(result.Name)
	nameEntry.ActionItem = widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		lo.Must0(result.Name.Set(faker.Word()))
	})
	protocolLabel := widget.NewLabel("Protocol")
	protocolSelection := widget.NewSelect(protocol.AvailableProtocols, func(s string) {
		err := result.Protocol.Set(s)
		if err != nil {
			panic(err)
		}
	})
	protocolSelection.SetSelectedIndex(0)
	addressLabel := widget.NewLabel("Address")
	addressEntry := widget.NewEntryWithData(result.Address)
	addressEntry.Validator = validation.NewAllStrings(func(s string) error {
		if s == "" {
			return errors.New("value is empty")
		}
		return nil
	})
	portLabel := widget.NewLabel("Port")
	portEntry := component.NewNumericalEntryWithData(result.Port)
	grid := container.New(layout.NewFormLayout(),
		nameLabel,
		nameEntry,
		protocolLabel,
		protocolSelection,
		addressLabel,
		addressEntry,
		portLabel,
		portEntry,
	)
	form := &ConnectionForm{
		selection:    protocolSelection,
		AddressEntry: addressEntry,
		PortEntry:    portEntry,
	}

	return grid, result, form
}

func NewAdvanceOption() *widget.Accordion {
	acceptor := widget.NewAccordionItem("AcceptorOption",
		widget.NewCheckGroup([]string{
			"Save To File",
			"Wrap",
		}, func(strings []string) {
			fmt.Println(strings)
		}),
	)
	sender := widget.NewAccordionItem("SenderOption",
		widget.NewCheckGroup([]string{
			"Base",
			"ScheduleTime",
		}, func(strings []string) {
			fmt.Println(strings)
		}),
	)
	return widget.NewAccordion(acceptor, sender)
}

func newButtons(
	onCancel func(),
	onSave func(result ConnectionDialogResult),
) *fyne.Container {
	var confirm = widget.NewButtonWithIcon("OK", theme.ConfirmIcon(), func() {
		onSave(ConnectionDialogResult{})
	})
	confirm.Importance = widget.HighImportance
	buttons := container.NewGridWithColumns(
		2,
		widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {
			onCancel()
		}),
		confirm)
	return buttons
}

func NewCreateConnectionForm(parameter ConnectionFormParameter) (*fyne.Container, *ConnectionForm, ConnectionDialogResult) {
	var requireForm, result, form = requiresForm()
	var testConnectionBtn = widget.NewButtonWithIcon("TestConnection", theme.InfoIcon(), func() {
		parameter.OnSave(result)
	})
	testConnectionBtn.Importance = widget.WarningImportance
	c := container.NewVBox(
		requireForm,
		NewAdvanceOption(),
		testConnectionBtn,
		newButtons(parameter.OnCancel, parameter.OnSave),
	)
	return c, form, result
}
