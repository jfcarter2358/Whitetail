package AST

import (
    "strings"
    // "log"
	"regexp"
	"whitetail/index"
	"github.com/google/uuid"
	"strconv"
	"sort"
	"whitetail/logging"
)

type AST struct {
    Operation string
	Left      string
	Right     string
}

func Parse(query string) []string {
	// fix query string formatting
	lCleanup1 := regexp.MustCompile(`\((\S)`)
	lCleanup2 := regexp.MustCompile(`(\S)\(`)
	rCleanup1 := regexp.MustCompile(`(\S)\)`)
	rCleanup2 := regexp.MustCompile(`\)(\S)`)

	// replace all 'open paren -> text' with 'open paren -> space -> text'
	for lCleanup1.MatchString(query) {
		query = lCleanup1.ReplaceAllString(query, "( $1")
	}
	// replace all 'text -> open paren' with 'text -> space -> open paren'
	for lCleanup2.MatchString(query) {
		query = lCleanup2.ReplaceAllString(query, "$1 (")
	}
	// replace all `text -> close paren` with 'text -> space -> close paren'
	for rCleanup1.MatchString(query) {
		query = rCleanup1.ReplaceAllString(query, "$1 )")
	}
	// replace all `close paren -> text` with 'close paren -> space -> text'
	for rCleanup2.MatchString(query) {
		query = rCleanup2.ReplaceAllString(query, ") $1")
	}
	
	// add beginning and end braces to query
	query = "( " + query + " )"

	// create our map of trees to hold the structure of the query
	ASTs := make(map[string]AST)

	pattern := regexp.MustCompile(`\(\s(\S*)\s(AND|OR|NOT|XOR|LIMIT|ORDER_ASCEND|ORDER_DESCEND)\s(\S*)\s\)`)
	topAST := AST{}
	didMatch := false
	for true {
		groups := pattern.FindStringSubmatch(query)
		if groups == nil {
			break
		}
		didMatch = true
		newAST := AST{Operation: groups[2], Left: groups[1], Right: groups[3]}
		id := "%" + uuid.New().String() + "%"
		ASTs[id] = newAST
		topAST = newAST
		query = strings.Replace(query, groups[0], id, 1)
	}

	if didMatch == false {
		index, err := Index.GetIndexByKey(query[2:len(query) - 2]) 
        if err != nil {
            return []string{}
        }
		out := strings.Split(index.IDs, ",")
		return out
	}
	out := Operate(topAST, ASTs)
	return out
}

func Operate(self AST, ASTs map[string]AST) []string {
	l := []string{}
	r := []string{}
	NON_LOGIC_OPERATIONS := []string{"LIMIT", "ORDER_ASCEND", "ORDER_DESCEND"}
	match, _ := regexp.Match(`%\S*%`, []byte(self.Left))
	if match {
		l = Operate(ASTs[self.Left], ASTs)
	} else {
		index, err := Index.GetIndexByKey(self.Left) 
		if err == nil {
			l = strings.Split(index.IDs, ",")
		}
	}
	if contains(NON_LOGIC_OPERATIONS, self.Operation) == false {
		match, _  = regexp.Match(`%\S*%`, []byte(self.Right))
		if match {
			r = Operate(ASTs[self.Right], ASTs)
		} else {
			index, err := Index.GetIndexByKey(self.Right) 
			if err == nil {
				r = strings.Split(index.IDs, ",")
			}
		}

		if self.Operation == "AND" {
			return AND(l, r)
		}
		if self.Operation == "OR" {
			return OR(l, r)
		}
		if self.Operation == "NOT" {
			return NOT(l, r)
		}
		if self.Operation == "XOR" {
			return XOR(l, r)
		}
		return []string{}
	} else {
		if self.Operation == "LIMIT" {
			return LIMIT(l, self.Right)
		}
		if self.Operation == "ORDER_ASCEND" {
			return ORDER_ASCEND(l, self.Right)
		}
		if self.Operation == "ORDER_DESCEND" {
			return ORDER_DESCEND(l, self.Right)
		}
		return []string{}
	}
}

func AND(l, r []string) []string {
	out := []string{}
	for _, i := range l {
		if contains(r, i) {
			out = append(out, i)
		}
	}
	return out
}

func OR(l, r []string) []string {
	out := l
	for _, i := range(r) {
		if contains(out, i) == false {
			out = append(out, i)
		}
	}
	return out
}

func NOT(l, r []string) []string {
	out := []string{}
	for _, i := range l {
		if contains(r, i) == false {
			out = append(out, i)
		}
	}
	return out
}

func XOR(l, r []string) []string {
	out := []string{}
	for _, i := range l {
		if contains(r, i) {
			out = append(out, i)
		}
	}
	for _, i := range r {
		if contains(l, i) {
			out = append(out, i)
		}
	}
	return out
}

func LIMIT(l []string, r string) []string {
	limit, err := strconv.Atoi(r)
	if err != nil {
		return []string{}
	}
	if limit < len(l) {
		return l[:limit]
	}
	return l
}

func ORDER_ASCEND(l []string, r string) []string {
	logs := []Logging.Log{}
	for _, id := range(l) {
		log, err := Logging.GetLogByID(id)
		if err == nil {
			logs = append(logs, *log)
		}
	}
	if r == "Text" {
		sort.Slice(logs[:], func(i, j int) bool {
			return logs[i].Text > logs[j].Text
		})
		ids := []string{}
		for _, log := range logs {
			ids = append(ids, log.ID)
		}
		return ids
	}
	if r == "Timestamp" {
		sort.Slice(logs[:], func(i, j int) bool {
			return logs[i].Timestamp > logs[j].Timestamp
		})
		ids := []string{}
		for _, log := range logs {
			ids = append(ids, log.ID)
		}
		return ids
	}
	if r == "Level" {
		sort.Slice(logs[:], func(i, j int) bool {
			return logs[i].Level > logs[j].Level
		})
		ids := []string{}
		for _, log := range logs {
			ids = append(ids, log.ID)
		}
		return ids
	}
	if r == "Service" {
		sort.Slice(logs[:], func(i, j int) bool {
			return logs[i].Service > logs[j].Service
		})
		ids := []string{}
		for _, log := range logs {
			ids = append(ids, log.ID)
		}
		return ids
	}
	return []string{}
}

func ORDER_DESCEND(l []string, r string) []string {
	logs := []Logging.Log{}
	for _, id := range(l) {
		log, err := Logging.GetLogByID(id)
		if err == nil {
			logs = append(logs, *log)
		}
	}
	if r == "Text" {
		sort.Slice(logs[:], func(i, j int) bool {
			return logs[i].Text < logs[j].Text
		})
		ids := []string{}
		for _, log := range logs {
			ids = append(ids, log.ID)
		}
		return ids
	}
	if r == "Timestamp" {
		sort.Slice(logs[:], func(i, j int) bool {
			return logs[i].Timestamp < logs[j].Timestamp
		})
		ids := []string{}
		for _, log := range logs {
			ids = append(ids, log.ID)
		}
		return ids
	}
	if r == "Level" {
		sort.Slice(logs[:], func(i, j int) bool {
			return logs[i].Level < logs[j].Level
		})
		ids := []string{}
		for _, log := range logs {
			ids = append(ids, log.ID)
		}
		return ids
	}
	if r == "Service" {
		sort.Slice(logs[:], func(i, j int) bool {
			return logs[i].Service < logs[j].Service
		})
		ids := []string{}
		for _, log := range logs {
			ids = append(ids, log.ID)
		}
		return ids
	}
	return []string{}
}

func contains(slice []string, value string) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}
	return false
}