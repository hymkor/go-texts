[![GoDoc](https://godoc.org/github.com/zetamatta/go-texts?status.svg)](https://godoc.org/github.com/zetamatta/go-texts)

go-texts
========

[go-texts/mbcs](./mbcs)
-----------------------
### translate string between ANSI and UTF8

	ansi, err := mbcs.UtoA("UTF8文字列", mbcs.ConsoleCP(), true)
	if err != nil {
		return err
	}

	utf8, err := mbcs.AtoU(ansi, mbcs.ConsoleCP())
	if err != nil {
		return err
	}
	fmt.Printf("Ok: %s\n", utf8)
	return nil

### reader converting from ANSI,UTF8 or UTF16 to UTF8

	sc := bufio.NewScanner(mbcs.NewAutoDetectReader(os.Stdin, mbcs.ConsoleCP()))
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

- [go-texts/filter](./filter) - filtering class which converts line by line as [transform](https://godoc.org/golang.org/x/text/transform)
- ByteOrderMark - `[]byte{0xEF, 0xBB, 0xBF}`
- SortedKey - it makes sorted strings' array from keys of the given map whose key's type is string.
