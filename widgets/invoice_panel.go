package widgets

import (
	"github.com/therecipe/qt/widgets"
	"schannel-qt5/parser"
	"sort"
)

// InvoicePanel 显示账单状态
type InvoicePanel struct {
	widgets.QWidget

	status       *ColorLabel
	showInvoices *widgets.QPushButton

	// 缓存的账单信息
	invoices []*parser.Invoice
}

// sortInvoices 将账单按开始日期排序
func (panel *InvoicePanel) sortInvoices() {
	sort.Slice(panel.invoices, func(i, j int) bool {
		if panel.invoices[i].StartDate.Before(panel.invoices[j].StartDate) {
			return true
		}

		return false
	})
}

// NewInvoicePanelWithData 生成InvoicePanel
func NewInvoicePanelWithData(data []*parser.Invoice) *InvoicePanel {
	panel := NewInvoicePanel(nil, 0)
	panel.invoices = make([]*parser.Invoice, 0, len(data))
	copy(panel.invoices, data)
	panel.sortInvoices()

	group := widgets.NewQGroupBox2("账单情况", nil)
	hbox := widgets.NewQHBoxLayout()

	text, isPaid := panel.invoices[len(panel.invoices)-1].GetStatus()
	if isPaid {
		panel.status = NewColorLabelWithColor(text, "green")
	} else {
		panel.status = NewColorLabelWithColor(text, "red")
	}
	hbox.AddWidget(panel.status, 0, 0)

	panel.showInvoices = widgets.NewQPushButton2("详细账单", nil)
	panel.showInvoices.ConnectClicked(panel.showInvoiceDialog)
	hbox.AddWidget(panel.showInvoices, 0, 0)

	group.SetLayout(hbox)
	mainLayout := widgets.NewQHBoxLayout()
	mainLayout.AddWidget(group, 0, 0)
	panel.SetLayout(mainLayout)

	return panel
}

// showInvoiceDialog 显示详细信息对话框
func (panel *InvoicePanel) showInvoiceDialog(_ bool) {
	dialog := NewInvoiceDialogWithData(panel.invoices)
	dialog.Exec()
}