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

func (u *Update) GetRelatedRules(rules []*PageRule) []*PageRule {
	relatedRules := make([]*PageRule, 0)
	for _, rule := range rules {
		rule1HasPages := slices.Contains(u.Pages, rule.Rule[0])
		rule2HasPages := slices.Contains(u.Pages, rule.Rule[1])
		if rule1HasPages && rule2HasPages {
			relatedRules = append(relatedRules, rule)
		}
	}
	return relatedRules
}

func (u *Update) IsCorrect(rules []*PageRule) bool {
	fmt.Println("Testing", u.Pages)
	for _, rule := range rules {
		fmt.Println("Rule", rule)
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
		relatedRules := update.GetRelatedRules(p.PageRules)
		if update.IsCorrect(relatedRules) {
			fmt.Println("CORRECT")
			correctUpdates = append(correctUpdates, update)
		} else {
			fmt.Println("NO")
		}
	}
	return correctUpdates
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
	res, _ := os.ReadFile("inputtest.txt")
	stringRes := strings.Split(string(res), "\n\n")
	printer := NewPrinter(stringRes[0], stringRes[1])
	fmt.Println(printer.GetCorrectUpdates())
}
