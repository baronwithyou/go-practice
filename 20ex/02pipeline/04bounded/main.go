package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

const n = 5

// 在03parallel的基础上进行修改 - 因为为每个文件开启一个goroutine太消耗内存了
// 使用固定个数的goroutine进行消费

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

// 读取文件写到channel中
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}

			return nil
		})
	}()

	return paths, errc
}

// 将文件中的信息写入到result channel中
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result{path, md5.Sum(data), err}:
		case <-done:
			return
		}
	}
}

func MD5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFiles(done, root)
	// 下面的代码是有问题的，因为此处的 `<-errc` 会将进程block
	// 因为还没到消费阶段，walkFiles中的paths的写入也会block
	//if err := <-errc; err != nil {
	//	return nil, err
	//}

	var wg sync.WaitGroup
	c := make(chan result)

	wg.Add(n)
	for i := 0; i < n; i++ {
		// 不使用 `go digester(done, paths, c)` 的形式污染digester方法
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}

	go func() {
		defer close(c)
		wg.Wait()
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	if err := <-errc; err != nil {
		return nil, err
	}

	return m, nil
}

func main() {
	m, err := MD5All(os.Args[1])
	if err != nil {
		panic(err)
	}

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)

	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
