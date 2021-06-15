package generate

import (
	"context"
	"fmt"
	"strings"
)

func getAPIInjectFileName(dir string) string {
	fullname := fmt.Sprintf("%s/internal/app/api/main.go", dir)
	return fullname
}

func insertAPIInject(ctx context.Context, dir, name string) error {
	injectContent := fmt.Sprintf("%sSet,", name)
	injectStart := 0
	insertFn := func(line string) (data string, flag int, ok bool) {
		if injectStart == 0 && strings.Contains(line, "var APISet = wire.NewSet(") {
			injectStart = 1
			return
		}

		if injectStart == 1 && strings.Contains(line, ")") {
			injectStart = -1
			data = injectContent
			flag = -1
			ok = true
			return
		}

		return "", 0, false
	}

	filename := getAPIInjectFileName(dir)
	err := insertContent(filename, insertFn)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", filename)

	return execGoFmt(filename)
}
