/*Package app 使用 goaccess 外部命令来分析日志文件并生成 html 和 json
 */
package app

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"path"
	"time"
)

// CreateOutputFile 运行 goaccess 生成 json 文件, gzips 是可以用 gzip 读取的文件
// 使用 gunzip 命去读取输入文件，生成 json 和 html 文件
func CreateOutputFile(zipfiles []string, dest string) error {
	for _, f := range zipfiles {
		log.Printf("analysis file: %s\n", f)
	}

	outPath := path.Clean(dest)
	log.Printf("the outpath is : %s\n", outPath)

	res, err := readLogFile(zipfiles)
	if err != nil {
		return fmt.Errorf("read gunzip command output failure. %w", err)
	}

	// write the gzip command output to goaccess stdin
	json := fmt.Sprintf("%s.json", time.Now().Format("20060102"))
	html := fmt.Sprintf("%s.html", time.Now().Format("20060102"))

	// 调用 `goaccess` 外部命令
	gaCmd := exec.Command(
		"goaccess",
		"--log-format=COMBINED",
		// `--date-format='%d%m%Y'`,
		"--output", path.Join(outPath, json),
		"--output", path.Join(outPath, html),
		"-", // 这个参数一定要加。表示个管道读取数据。不加的的话，在使用 `crontab` 运行程序会出错
	)

	log.Printf("run cmd: %s", gaCmd) // record log

	gaPipe, err := gaCmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("create goaccess input pipe failure. %w", err)
	}

	// 这里必须在一个 goroutine 里面执行，不然程序的运行会卡住
	go func() {
		defer gaPipe.Close()
		io.WriteString(gaPipe, string(res))
	}()

	_, err = gaCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run goaccess failure. %w", err)
	}
	return nil
}

// readLogFile use gunzip external command to read gzip files
// return `[]byte`
func readLogFile(files []string) ([]byte, error) {
	// use gnuzip read archive files
	// read the result from std output pipe

	// 拼装执行命令的参数
	cmdStr := []string{"-c"}
	for _, f := range files {
		cmdStr = append(cmdStr, f)
	}

	gzipCmd := exec.Command("gunzip", cmdStr...)
	log.Printf("run cmd: %s\n", gzipCmd)

	gzipOutPipe, err := gzipCmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("create gunzip output pipe failure. %w", err)
	}

	if err = gzipCmd.Start(); err != nil {
		return nil, fmt.Errorf("run gzip command failure. %w", err)
	}

	return ioutil.ReadAll(gzipOutPipe)
}
