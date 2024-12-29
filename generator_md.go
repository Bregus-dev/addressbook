package main

import (
	"fmt"
	"strings"
)

// //////////////////////////////////////////////////////////////

func MdHeader(level int, text string) string {
	if level < 1 || level > 6 {
		level = 1
	}
	return fmt.Sprintf("%s %s\n\n", strings.Repeat("#", level), text)
}

func MdList(items []string) string {
	var result strings.Builder
	for _, item := range items {
		result.WriteString(fmt.Sprintf(" - %s\n", item))
	}
	result.WriteString("\n")
	return result.String()
}

func MdOrderedList(items []string) string {
	var result strings.Builder
	for i, item := range items {
		result.WriteString(fmt.Sprintf("%d. %s\n", i+1, item))
	}
	result.WriteString("\n")
	return result.String()
}

func MdTable(headers []string, rows [][]string) string {
	var result strings.Builder

	result.WriteString("| " + strings.Join(headers, " | ") + " |\n")
	result.WriteString("|" + strings.Repeat("---|", len(headers)) + "\n")

	for _, row := range rows {
		result.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}

	result.WriteString("\n")
	return result.String()
}
