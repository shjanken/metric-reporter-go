/*Package app 使用 goaccess 外部命令来分析日志文件并生成 html 和 json
 */
package app

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"path"
)

// CreateJSONFile 运行 goaccess 生成 json 文件, 使用 files 作为输入文件，生成的文件存放到 dest
func CreateJSONFile(files []string, dest string) error {
	outPath := path.Dir(path.Clean(dest))
	res, err := readLogFile(files)
	if err != nil {
		return fmt.Errorf("read gunzip command output failure. %w", err)
	}

	// write the gzip command output to goaccess stdin
	jsonFile := fmt.Sprintf("%s/report.json", outPath)
	htmlFile := fmt.Sprintf("%s/report.html", outPath)
	gaCmd := exec.Command(
		"goaccess", "--log-format=COMBINED",
		"-o", jsonFile,
		"-o", htmlFile)
	gaPipe, err := gaCmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("create goaccess input pipe failure. %w", err)
	}
	go func() {
		defer gaPipe.Close()
		io.WriteString(gaPipe, string(res))
	}()

	_, err = gaCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run goaccess failure. %w", err)
	} /*  */
	return nil
}

func readLogFile(files []string) ([]byte, error) {
	// use gnuzip read archive files
	// read the result from std output pipe

	// 拼装执行命令的参数
	cmdStr := []string{"-c"}
	for _, f := range files {
		cmdStr = append(cmdStr, f)
	}

	gzipCmd := exec.Command("gunzip", cmdStr...)
	gzipOutPipe, err := gzipCmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("create gunzip output pipe failure. %w", err)
	}

	if err = gzipCmd.Start(); err != nil {
		return nil, fmt.Errorf("run gzip command failure. %w", err)
	}

	return ioutil.ReadAll(gzipOutPipe)
}
