package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func FindSymmetricPoems() [][]string {
	path := "./5.txt"
	xss := ReadList(path)
	dict := CreateDictionary(xss)
	results := ParallelSearch(dict)
	return RunesToStrings(results)
}

// ReadList
func ReadList(path string) []string {
	AbsolutePath := func(_path interface{}) (interface{}, error) {
		var path string = _path.(string)
		return filepath.Abs(path)
	}

	OpenFile := func(_path interface{}) (interface{}, error) {
		var path = _path.(string)
		return os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	}

	Scan := func(_file interface{}) (interface{}, error) {
		var file *os.File = _file.(*os.File)
		var result = []string{}
		sc := bufio.NewScanner(file)
		for sc.Scan() {
			result = append(result, sc.Text())
		}
		defer file.Close()

		err := sc.Err()
		if err != nil {
			fmt.Println(err)
			return []string{}, err
		} else {
			return result, nil
		}
	}

	_ReadList := func(path string) []string {
		list, err := Bind(AbsolutePath, OpenFile, Scan)(path)
		if err != nil {
			fmt.Println(err)
			return []string{}
		} else {
			return list.([]string)
		}
	}
	return _ReadList(path)
}

func Bind(fs ...func(interface{}) (interface{}, error)) func(interface{}) (interface{}, error) {
	return func(x interface{}) (interface{}, error) {
		data := x
		var err error
		for _, f := range fs {
			data, err = f(data)
			if err != nil {
				return nil, err
			}
		}
		return data, nil
	}
}

// ParallelSearch
func ParallelSearch(dict map[rune][][]rune) [][][]rune {
	var results = [][][]rune{}
	var resultsPtr interface{} = &results
	ParallelForEach(Search(dict), dict, resultsPtr)
	return results
}

func ParallelForEach(worker func(interface{}, interface{}, interface{}), dict map[rune][][]rune, results interface{}) {
	var wg sync.WaitGroup
	wg.Add(len(dict))

	f := func(key interface{}, value interface{}) {
		defer wg.Done()
		worker(key, value, results)
	}

	for key, value := range dict {
		go f(key, value)
	}
	wg.Wait()
}

// Search
func Search(dict map[rune][][]rune) func(key interface{}, _lines interface{}, _results interface{}) {
	return func(key interface{}, _lines interface{}, _results interface{}) {
		results := _results.(*[][][]rune)
		lines := _lines.([][]rune)
		for _, line := range lines {
			var result [][]rune = [][]rune{line}
			DFS(dict, result, results, 1)
		}
	}
}

func DFS(dict map[rune][][]rune, result [][]rune, results *[][][]rune, index int) {
	if FoundQ(index) {
		*results = append(*results, result)
		PrintResult(result)
	} else {
		_DFS(dict, result, results, index)
	}
}

func _DFS(dict map[rune][][]rune, result [][]rune, results *[][][]rune, index int) {
	nextCharacter := result[0][index]
	lines, ok := dict[nextCharacter]
	if !ok {
		return
	}

	for _, line := range lines {
		if Validate(index, result, line) {
			DFS(dict, append(result, line), results, index+1)
		}
	}
}

func FoundQ(index int) bool {
	return index == 5
}

func PrintResult(result [][]rune) {
	for _, x := range result {
		fmt.Println(string(x))
	}
	fmt.Println("")
}

func Validate(index int, result [][]rune, rs []rune) bool {
	for i := 1; i < index; i++ {
		if result[i][index] != rs[i] {
			return false
		}
	}
	return true
}

// RunesToStrings
func RunesToStrings(rsss [][][]rune) [][]string {
	var sss = make([][]string, len(rsss))
	for index, rss := range rsss {
		var ss = []string{}
		for _, rs := range rss {
			ss = append(ss, string(rs))
		}
		sss[index] = ss
	}
	return sss
}

// CreateDictionary
func CreateDictionary(xss []string) map[rune][][]rune {
	var dict = map[rune][][]rune{}
	for _, xs := range xss {
		AppendLine(dict, xs)
	}
	return dict
}

func AppendLine(dict map[rune][][]rune, xs string) {
	rs := []rune(xs)
	key := rs[0]
	rss, ok := dict[key]
	if !ok {
		dict[key] = [][]rune{rs}
	} else {
		dict[key] = append(rss, rs)
	}
}

func main() {
	var results = FindSymmetricPoems()
	fmt.Println(results)
}
