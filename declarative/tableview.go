// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type TableView struct {
	Widget                     **walk.TableView
	Name                       string
	StretchFactor              int
	Row                        int
	RowSpan                    int
	Column                     int
	ColumnSpan                 int
	Font                       Font
	Model                      walk.TableModel
	AlternatingRowBGColor      walk.Color
	CheckBoxes                 bool
	ItemStateChangedEventDelay int
	LastColumnStretched        bool
	ReorderColumnsEnabled      bool
	SingleItemSelection        bool
	OnCurrentIndexChanged      walk.EventHandler
	OnSelectedIndexesChanged   walk.EventHandler
	OnItemActivated            walk.EventHandler
}

func (tv TableView) Create(parent walk.Container) error {
	w, err := walk.NewTableView(parent)
	if err != nil {
		return err
	}

	return InitWidget(tv, w, func() error {
		if err := w.SetModel(tv.Model); err != nil {
			return err
		}

		w.SetAlternatingRowBGColor(tv.AlternatingRowBGColor)
		w.SetCheckBoxes(tv.CheckBoxes)
		w.SetItemStateChangedEventDelay(tv.ItemStateChangedEventDelay)
		if err := w.SetLastColumnStretched(tv.LastColumnStretched); err != nil {
			return err
		}
		w.SetReorderColumnsEnabled(tv.ReorderColumnsEnabled)
		if err := w.SetSingleItemSelection(tv.SingleItemSelection); err != nil {
			return err
		}

		if tv.OnCurrentIndexChanged != nil {
			w.CurrentIndexChanged().Attach(tv.OnCurrentIndexChanged)
		}
		if tv.OnSelectedIndexesChanged != nil {
			w.SelectedIndexesChanged().Attach(tv.OnSelectedIndexesChanged)
		}
		if tv.OnItemActivated != nil {
			w.ItemActivated().Attach(tv.OnItemActivated)
		}

		if tv.Widget != nil {
			*tv.Widget = w
		}

		return nil
	})
}

func (tv TableView) CommonInfo() (name string, stretchFactor, row, rowSpan, column, columnSpan int) {
	return tv.Name, tv.StretchFactor, tv.Row, tv.RowSpan, tv.Column, tv.ColumnSpan
}

func (tv TableView) Font_() *Font {
	return &tv.Font
}
