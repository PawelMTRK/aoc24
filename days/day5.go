package days

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type PageRule struct {
	Rule [2]int
}

func NewPageRule(str string) PageRule {
	rule := strings.Split(str, "|")
	r1, _ := strconv.Atoi(rule[0])
	r2, _ := strconv.Atoi(rule[1])
	return PageRule{Rule: [2]int{r1, r2}}
}

type Update struct {
	Pages []int
}

func NewUpdate(str string) Update {
	pagesSplit := strings.Split(str, ",")
	pages := make([]int, 0)
	for _, p := range pagesSplit {
		n, _ := strconv.Atoi(p)
		pages = append(pages, n)
	}
	return Update{Pages: pages}
}

func (u *Update) IsCorrect(rules []*PageRule) bool {
	for _, rule := range rules {
		// skip checking the rule if it doesn't
		// contain processed pages
		rule1HasPages := slices.Contains(u.Pages, rule.Rule[0])
		rule2HasPages := slices.Contains(u.Pages, rule.Rule[1])
		if !(rule1HasPages && rule2HasPages) {
			continue
		}
		// check if the rule applies
		p1Index := slices.Index(u.Pages, rule.Rule[0])
		p2Index := slices.Index(u.Pages, rule.Rule[1])
		if p1Index > p2Index {
			return false
		}
	}
	return true
}

type Printer struct {
	PageRules []*PageRule
	Updates   []Update
}

func (p *Printer) GetCorrectUpdates() []Update {
	correctUpdates := make([]Update, 0)
	for _, update := range p.Updates {
		if update.IsCorrect(p.PageRules) {
			correctUpdates = append(correctUpdates, update)
		}
	}
	return correctUpdates
}

func (p *Printer) GetSum() int {
	total := 0
	for _, v := range p.GetCorrectUpdates() {
		middlePage := v.Pages[len(v.Pages)/2]
		total += middlePage
	}
	return total
}

func NewPrinter(rulesStr, updateStr string) Printer {
	pageRules := make([]*PageRule, 0)
	updates := make([]Update, 0)
	for _, r := range strings.Split(rulesStr, "\n") {
		rule := NewPageRule(r)
		pageRules = append(pageRules, &rule)
	}
	for _, u := range strings.Split(updateStr, "\n") {
		updates = append(updates, NewUpdate(u))
	}
	return Printer{PageRules: pageRules, Updates: updates}
}

func Day5() {
	res, _ := os.ReadFile("input5.txt")
	stringRes := strings.Split(string(res), "\n\n")
	printer := NewPrinter(stringRes[0], stringRes[1])
	fmt.Println("Middle page numbers:", printer.GetSum())
}
