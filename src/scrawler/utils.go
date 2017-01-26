package scrawler

import (
  "fmt"
  "regexp"
  "strconv"
  "strings"
  "os"
  "bufio"
  "path/filepath"
)

func writeMaptoFile(m map[string]string, filePath string) error {
  f, err := os.Create(filePath)
  if err != nil {
    fmt.Printf("create map file error: %v\n", err)
    return err
  }
  defer f.Close()

  w := bufio.NewWriter(f)
  for k, v := range m {
    lineStr := fmt.Sprintf("%s^%s", k, v)
    fmt.Fprintln(w, lineStr)
  }
  return w.Flush()
}

func getPageData(filePath, startUrl, loginCookies string) error {
  f, err := os.Create(filePath)
  if err != nil {
          fmt.Printf("create map file error: %v\n", err)
          return err
  }
  defer f.Close()

  w := bufio.NewWriter(f)
  //for k, v := range m {
  mstartResp, _ := DoRequest(`GET`, startUrl, ``, loginCookies, ``, header)
  reg := regexp.MustCompile(`<span class="ctt">(.*?)</span>`)
  arrStart := reg.FindAllStringSubmatch(mstartResp, -1)
  if len(arrStart) > 0 {
    fmt.Println(len(arrStart))
    for i := 0; i < len(arrStart); i++ {
       //去除所有尖括号内的HTML代码，并换成换行符
      re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
      arrStart[i][1] = re.ReplaceAllString(arrStart[i][1], "")
      fmt.Println(arrStart[i][1])
      fmt.Fprintln(w, arrStart[i][1])
    }
    //}
    //lineStr := fmt.Sprintf("%s^%s", k, v)
    //fmt.Fprintln(w, lineStr)
  }
  return w.Flush()
}

func getFilelist(path string) {
  i := 0
  err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
    if ( f == nil ) {return err}
    if f.IsDir() {return nil}
    i++
    fmt.Println(path)
    err = os.Rename(path, strconv.Itoa(i))     //重命名 C:\\log\\2013.log 文件为install.txt
    if err != nil {
        //如果重命名文件失败,则输出错误 file rename Error!
        fmt.Println("file rename Error!")
        //打印错误详细信息
        fmt.Printf("%s", err)
    } else {
        //如果文件重命名成功,则输出 file rename OK!
        fmt.Println("file rename OK!")
    }

    return nil
  })
  if err != nil {
    fmt.Printf("filepath.Walk() returned %v\n", err)
  }
}

func ReadLine(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
  var urls
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
    append(urls, line)
		if err != nil {
			if err == io.EOF {
				return urls, nil
			}
			return nil, err
		}
	}
	return urls, nil
}
