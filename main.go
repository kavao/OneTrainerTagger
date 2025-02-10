package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/ncruces/zenity"
)

type Config struct {
	suffixText  string
	selectedDir string
}

type AppState struct {
	config       *Config
	theme        *material.Theme
	selectDirBtn widget.Clickable
	processBtn   widget.Clickable
	suffixEditor widget.Editor
	status       string
}

func main() {
	go func() {
		w := app.NewWindow(app.Title("FineTuning用元データタグ付け補助 for OneTrainer"))
		w.Option(app.Size(unit.Dp(800), unit.Dp(600)))

		state := &AppState{
			config: &Config{},
			theme:  material.NewTheme(),
		}
		state.suffixEditor.SetText(", best quality, masterpiece")

		var ops op.Ops
		for e := range w.Events() {
			switch e := e.(type) {
			case system.DestroyEvent:
				os.Exit(0)
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				state.layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func (state *AppState) layout(gtx layout.Context) layout.Dimensions {
	// ボタンのクリックハンドラを修正
	if state.selectDirBtn.Clicked() {
		if dir, err := zenity.SelectFile(
			zenity.Title("ディレクトリを選択"),
			zenity.Directory(),
		); err == nil {
			state.config.selectedDir = dir
			state.status = "選択されたディレクトリ: " + dir
		}
	}

	if state.processBtn.Clicked() && state.config.selectedDir != "" {
		state.config.suffixText = state.suffixEditor.Text()
		processFiles(state.config)
		state.status = "処理が完了しました"
	}

	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceEvenly,
	}.Layout(gtx,
		// 左側の説明列（幅を固定）
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max.X = gtx.Dp(400)
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(layout.Spacer{Height: 20}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Body1(state.theme, "1. 画像ファイルが含まれるディレクトリを選択してください").Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: 40}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Body1(state.theme, "2. 画像の説明文の末尾に追加するタグを入力してください").Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: 40}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Body1(state.theme, "3. 処理を実行すると、画像と同じ名前のtxtファイルが作成されます").Layout(gtx)
				}),
			)
		}),
		// 右側のコントロール列
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max.X = gtx.Dp(300)
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(layout.Spacer{Height: 20}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(state.theme, &state.selectDirBtn, "ディレクトリを選択").Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: 20}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Editor(state.theme, &state.suffixEditor, "追加するタグ").Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: 20}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(state.theme, &state.processBtn, "処理実行").Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: 20}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Body1(state.theme, state.status).Layout(gtx)
				}),
			)
		}),
	)
}

func processFiles(config *Config) {
	filepath.Walk(config.selectedDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if isImageFile(info.Name()) {
			baseName := filepath.Base(path)
			ext := filepath.Ext(baseName)
			nameWithoutExt := strings.TrimSuffix(baseName, ext)

			cleanName := removeSerialNumber(nameWithoutExt)
			commaName := strings.ReplaceAll(cleanName, "_", ", ")

			txtContent := fmt.Sprintf("%s%s", commaName, config.suffixText)
			txtPath := filepath.Join(filepath.Dir(path), nameWithoutExt+".txt")

			err := os.WriteFile(txtPath, []byte(txtContent), 0644)
			if err != nil {
				fmt.Printf("エラー: %v\n", err)
			}
		}
		return nil
	})
}

func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}

func removeSerialNumber(name string) string {
	// UUID/ハッシュパターン（8-4-4-4-12形式）を削除
	re1 := regexp.MustCompile(`_[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)
	name = re1.ReplaceAllString(name, "")

	// 末尾の_数字を削除
	re2 := regexp.MustCompile(`_\d+$`)
	name = re2.ReplaceAllString(name, "")

	// 連続する_を1つにまとめる
	re3 := regexp.MustCompile(`_+`)
	return re3.ReplaceAllString(name, "_")
}
