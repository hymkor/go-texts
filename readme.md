[![GoDoc](https://godoc.org/github.com/zetamatta/go-texts?status.svg)](https://godoc.org/github.com/zetamatta/go-texts)

go-texts is the utility package for text-data

"go-texts"
=========

SortedKeys
---------
It makes sorted strings' array from keys of the given map whose key's type is string.

	map1 := map[string]string{
		"A": "alpha",
		"B": "beta",
		"C": "gamma",
	}

	for _, key1 := range texts.SortedKeys(map1) {
		fmt.Printf("%s: %s\n", key1, map1[key1])
	}


"go-texts/mbcs"
===============

Translate string between ANSI and UTF8
--------------------------------------

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

Reader converting from ANSI,UTF8 or UTF16 to UTF8
-------------------------------------------------

	sc := bufio.NewScanner(mbcs.NewAutoDetectReader(os.Stdin, mbcs.ConsoleCP()))
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

"go-texts/dos"
=====================

Call CMD.exe without troubles about double-quotation
----------------------------------------------------

	dos.System(`echo "ahaha" "ihihi" "ufufu"`)
