package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/oklookat/kp2imdb/cmd"
	"github.com/oklookat/kp2imdb/export"
	"github.com/oklookat/kp2imdb/imdb"

	_ "embed"
)

const (
	_linksPath = "./links.json"
)

var (
	//go:embed HELP.txt
	help string
)

func main() {
	var useManualLink bool
	flag.BoolVar(&useManualLink, "m", false, "Вручную вводить IMDB id при ошибках или ненаходе тайтла.")
	flag.Parse()

	kpFile := flag.Arg(0)
	if len(kpFile) == 0 {
		println(help)
		bufio.NewReader(os.Stdin).ReadString('\n')
		os.Exit(0)
		return
	}

	links, err := export.LoadLinks(_linksPath)
	chk(err)

	exported, err := export.LoadKinopoisk(kpFile)
	chk(err)

	stack := cmd.NewStack(50)
	barUid := stack.AddAlwaysBottom("")
	setBar := func(i int) {
		stack.AlwaysBottom[barUid] = fmt.Sprintf("\n⏳ %d/%d", i+1, len(exported))
		stack.Render()
	}

	var imdbIds []string
	for i, ke := range exported {
		setBar(i)

		imdbId, ok := links[ke.ID]
		if ok {
			// already linked.
			imdbIds = append(imdbIds, imdbId)
			continue
		}

		saveId := func(kpId, imdbId string) {
			imdbIds = append(imdbIds, imdbId)
			links[kpId] = imdbId
			err = export.SaveLinks(_linksPath, links)
			chk(err)
		}

		var found bool
		parsed := ke.ParsedName
		for i := 0; i < 2; i++ {
			searchMsg := func() string {
				return fmt.Sprintf("%s (%d)", parsed.Title, parsed.Year)
			}
			searchUid := stack.Add("🔍 " + searchMsg())
			stack.Render()

			imdbTitle, err := imdb.SearchTitle(parsed, i == 1)

			if err != nil {
				stack.Add("⚠️  " + err.Error())
				stack.Render()
			}

			if err != nil || imdbTitle == nil {
				stack.Stack[searchUid] = "❌ " + searchMsg()
				stack.Render()
				if i == 0 && ke.ParsedAltName != nil {
					parsed = *ke.ParsedAltName
					continue
				}
				break
			}

			stack.Stack[searchUid] = fmt.Sprintf("✅ %s ||| %v", searchMsg(), imdbTitle)
			stack.Render()
			saveId(ke.ID, imdbTitle.ID)
			found = true
			break
		}

		if !found && useManualLink {
			id, err := manualLink(&stack, &ke)
			chk(err)
			saveId(ke.ID, id)
		}
	}

	err = export.SaveIds(fmt.Sprintf("%d.txt", time.Now().Unix()), imdbIds)
	chk(err)
}

func chk(err error) {
	if err == nil {
		return
	}
	println("Ошибка: " + err.Error())
	println(`Возможные решения:
1. Убедитесь что у вас есть интернет.
2. Проверьте свой JSON файл.
3. Проверьте права доступа (links.json, создание файлов в текущей директории).
4. Проверьте links.json или удалите его.
Или опишите проблему тут: https://github.com/oklookat/kp2imdb/issues
`)
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func manualLink(st *cmd.Stack, ke *export.KpExport) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	st.Add(fmt.Sprintf(`🟦 Не найдено: %s | %s
🟦 Вставьте IMDB ID (например: tt6263850):`, ke.Name, ke.AltName))
	st.Render()
	id, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(id), err
}
