package main

import (
	"fmt"
	"log"

	"github.com/c-bata/go-prompt"
	"github.com/marcusolsson/tui-go"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "get", Description: "Command: Gets the query"},
		{Text: "nameStartsWith", Description: "Query: Gets characters with name"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

type mail struct {
	from    string
	subject string
	date    string
	body    string
}

var mails = []mail{
	{
		from:    "John Doe <john@doe.com>",
		subject: "Vacation pictures",
		date:    "Yesterday",
		body: `
Hey,
Where can I find the pictures from the diving trip?
Cheers,
John`,
	},
	{
		from:    "Jane Doe <jane@doe.com>",
		subject: "Meeting notes",
		date:    "Yesterday",
		body: `
Here are the notes from today's meeting.
/Jane`,
	},
	{
		from:    "Jane Doe <jane@doe.com>",
		subject: "Meeting notes",
		date:    "Yesterday",
		body: `
Here are the notes from today's meeting.
/Jane`,
	},
	{
		from:    "Jane Doe <jane@doe.com>",
		subject: "Meeting notes",
		date:    "Yesterday",
		body: `
Here are the notes from today's meeting.
/Jane`,
	},
	{
		from:    "Jane Doe <jane@doe.com>",
		subject: "Meeting notes",
		date:    "Yesterday",
		body: `
Here are the notes from today's meeting.
/Jane`,
	},
}

func main() {

	// data := get_data("characters?nameStartsWith=iron")
	// m, ok := gjson.Parse(data).Value().(map[string]interface{})

	// if ok {
	// 	fmt.Println(m["copyright"])
	// }
	setup()
	fmt.Println("Please select table.")
	t := prompt.Input("marvel> ", completer)
	fmt.Println("You selected " + t)
	// fmt.Printf("You choose %q\n", result)

	inbox := tui.NewTable(0, 0)
	inbox.SetColumnStretch(0, 3)
	inbox.SetColumnStretch(1, 2)
	inbox.SetColumnStretch(2, 1)
	inbox.SetFocused(true)

	for _, m := range mails {
		inbox.AppendRow(
			tui.NewLabel(m.subject),
			tui.NewLabel(m.from),
			tui.NewLabel(m.date),
		)
	}

	// var (
	// 	from    = tui.NewLabel("")
	// 	subject = tui.NewLabel("")
	// 	date    = tui.NewLabel("")
	// )

	// info := tui.NewGrid(0, 0)
	// info.AppendRow(tui.NewLabel("From:"), from)
	// info.AppendRow(tui.NewLabel("Subject:"), subject)
	// info.AppendRow(tui.NewLabel("Date:"), date)

	// body := tui.NewLabel("")
	// body.SetSizePolicy(tui.Preferred, tui.Expanding)

	// mail := tui.NewVBox(info, body)
	// mail.SetSizePolicy(tui.Preferred, tui.Expanding)

	// inbox.OnSelectionChanged(func(t *tui.Table) {
	// 	m := mails[t.Selected()]
	// 	from.SetText(m.from)
	// 	subject.SetText(m.subject)
	// 	date.SetText(m.date)
	// 	body.SetText(m.body)
	// })

	// // Select first mail on startup.
	// inbox.Select(0)

	root := tui.NewVBox(inbox, tui.NewLabel(""))

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Ctrl+Q", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
	setupCheck()
}
