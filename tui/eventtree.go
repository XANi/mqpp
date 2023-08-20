package tui

import (
	"fmt"
	"github.com/XANi/goneric"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"sort"
	"time"
)

var i = 0

func (t *TUI) fillTree(p []pterm.LeveledListItem, tree *EvTree, level int) []pterm.LeveledListItem {
	var label string
	if len(tree.lastMessage) > 0 {
		label = fmt.Sprintf("%s [%s]", tree.key, string(tree.lastMessage))
		if len(label) > 96 {
			label = label[0:96] + "..."
		}
	} else {
		label = tree.key
	}
	p = append(p, pterm.LeveledListItem{
		Level: level,
		Text:  label,
	})

	if tree == nil || tree.t == nil || len(tree.t) == 0 {
		return p
	}
	keys := goneric.MapSliceKey(tree.t)
	sort.Strings(keys)
	for _, k := range keys {
		v := tree.t[k]
		p = t.fillTree(p, v, level+1)
	}
	return p

}

func (t *TUI) EventTree() (string, error) {
	tl := []pterm.LeveledListItem{}
	tl = t.fillTree(tl, t.evTree, 0)
	tree, err := pterm.DefaultTree.WithRoot(putils.TreeFromLeveledList(tl)).Srender()
	if err != nil {
		return "", err
	}
	panels := pterm.Panels{
		{
			{
				Data: tree,
			},
			{
				Data: pterm.DefaultHeader.Sprintf("t: %s", time.Now().Format("15:04:05.00")),
			},
		},
		/*		{
					{
						Data: "This\npanel\ncontains\nmultiple\nlines",
					},
				},
				{
					{
						Data: pterm.Red("This is another\npanel line")},
					{Data: "This is the second panel\nwith a new line"},
				},*/
	}
	return pterm.DefaultPanel.WithPanels(panels).Srender()
}
