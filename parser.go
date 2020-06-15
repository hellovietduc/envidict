package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ParseFile(db *AVLTree) {
	file, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var curWord *Word
	var curDef *Definition
	var curDesc *Description

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			if curWord == nil {
				continue
			}

			if curDef != nil {
				if curDesc != nil {
					curDef.Descriptions = append(curDef.Descriptions, curDesc)
				}
				curWord.Definitions = append(curWord.Definitions, curDef)
			}
			db.Insert(curWord.Spelling, curWord)
			curWord = nil
			curDef = nil
			curDesc = nil
			continue
		}

		r := []rune(line)
		firstChar := string(r[0])
		rest := strings.TrimSpace(string(r[1:]))

		switch firstChar {
		case "@":
			tmp := strings.Split(rest, " ")
			if len(tmp) < 2 {
				break
			}

			curWord = &Word{
				Spelling:      tmp[0],
				Pronunciation: tmp[1],
				Definitions:   make([]*Definition, 0),
			}
		case "*":
			if curWord == nil {
				curWord = &Word{
					Definitions: make([]*Definition, 0),
				}
			}

			if curDef != nil {
				curWord.Definitions = append(curWord.Definitions, curDef)
			}

			curDef = &Definition{
				Kind:         rest,
				Descriptions: make([]*Description, 0),
			}
		case "-":
			if curDef == nil {
				curDef = &Definition{}
			}

			if curDesc != nil {
				curDef.Descriptions = append(curDef.Descriptions, curDesc)
			}

			curDesc = &Description{
				Meaning: rest,
			}
		case "=":
			if curDesc == nil {
				curDesc = &Description{}
			}
			curDesc.Example = strings.Replace(rest, "+ ", ": ", 1)
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
