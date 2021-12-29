package run

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func DownGTPic(dir, id string) (err error) {
	var (
		pic []string
	)
	pic, err = GetGT(id)
	if err != nil {
		return
	}
	if len(pic) == 0 {
		return errors.New("未匹配到图片")
	}
	err = downloadPic(dir, pic)
	if err != nil {
		return
	}
	return

}

// http://www.gtpso.com/index.php/home/index/viewTab?
func GetGT(id string) (pic []string, err error) {
	var (
		post *http.Request
	)

	post, err = http.NewRequest("GET", "http://www.gtpso.com/index.php/home/index/viewTab?id="+id, nil)
	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(post)
	if err != nil {
		return
	}
	defer res.Body.Close()
	buf := bufio.NewReader(res.Body)

	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		v := string(a)
		if strings.Contains(v, "img-fluid") {
			//     标题
			pic = ReadGTPic(v)
			return
		}

	}
	return
}

func ReadGTPic(v string) (pics []string) {
	s := strings.Index(v, "src=\"")
	e := strings.Index(v, "\" class=")
	if s == -1 || e == -1 {
		return nil
	}
	pics = append(pics, v[s+5:e])
	pics = append(pics, ReadGTPic(v[e+8:])...)
	return
}

func downloadPic(dir string, pics []string) (err error) {
	var (
		res *http.Response
	)
	for i, v := range pics {
		filesuffix := path.Ext(v)
		res, err = http.Get(v)
		if err != nil {
			return
		}
		defer res.Body.Close()

		if 0 != len(dir) {
			is_exist := false
			is_exist, err = PathExists(dir)
			if nil != err {
				return
			}
			if is_exist == false {
				//不存在文件夹时 先创建文件夹再上传
				err = os.MkdirAll(dir, os.ModePerm) //创建文件夹
				if err != nil {
					return
				}
			}
		}
		storeName := fmt.Sprintf("%d%s", i+1, filesuffix)
		f, err := os.OpenFile(dir+"/"+storeName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		//写入文件数据;
		io.Copy(f, res.Body)
		f.Close()
	}
	return
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
