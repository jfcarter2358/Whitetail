package AST

import (
    "strings"
    "log"
	"regexp"
	// "whitetail/index"
	"github.com/google/uuid"
	"strconv"
	// "sort"
	"whitetail/logging"
	"gorm.io/gorm"
)

type AST struct {
    Operation string
	Left      string
	Right     string
}

func Parse(query string) []Logging.Log {
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

	pattern := regexp.MustCompile(`\(\s(\S*\s?(?:=|>|>=|<=|<|!=|IN)\s?\S*|\S*)\s(AND|OR|NOT|XOR|LIMIT|ORDER_ASCEND|ORDER_DESCEND)\s(\S*\s?(?:=|>|>=|<=|<|!=|IN)\s?\S*|\S*)\s\)`)
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
		out := parseQuery(query[2:len(query) - 2])
		var logs []Logging.Log

		out.Find(&logs)
		return logs
	}
	out := Operate(topAST, ASTs)
	var logs []Logging.Log

	out.Find(&logs)
	return logs
}

func parseQuery(query string) *gorm.DB {
	pattern := regexp.MustCompile(`(\S*)\s?(=|>|>=|<=|<|!=|IN)\s?(\S*)`)
	groups := pattern.FindStringSubmatch(query)
	log.Println(query)
	log.Println(groups[1] + " " + groups[2] + " ?")
	log.Println(groups[3])
	if groups[2] == "IN" {
		Logging.DB.Where(groups[1] + " " + groups[2] + " ?", strings.Split(groups[3], ","))
	}
	return Logging.DB.Where(groups[1] + " " + groups[2] + " ?", groups[3])
}

func Operate(self AST, ASTs map[string]AST) *gorm.DB {
	var l *gorm.DB
	var r *gorm.DB
	NON_LOGIC_OPERATIONS := []string{"LIMIT", "ORDER_ASCEND", "ORDER_DESCEND"}
	match, _ := regexp.Match(`%\S*%`, []byte(self.Left))
	if match {
		l = Operate(ASTs[self.Left], ASTs)
	} else {
		l = parseQuery(self.Left)
	}
	if contains(NON_LOGIC_OPERATIONS, self.Operation) == false {
		match, _  = regexp.Match(`%\S*%`, []byte(self.Right))
		if match {
			r = Operate(ASTs[self.Right], ASTs)
		} else {
			r = parseQuery(self.Right)
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
		return nil
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
		return nil
	}
}

func AND(l, r *gorm.DB) *gorm.DB {
	return l.Where(r)
}

func OR(l, r *gorm.DB) *gorm.DB {
	return l.Or(r)
}

func NOT(l, r *gorm.DB) *gorm.DB {
	return l.Not(r)
}

func XOR(l, r *gorm.DB) *gorm.DB {
	// return l.Xor(r)
	return nil
}

func LIMIT(l *gorm.DB, r string) *gorm.DB {
	limit, err := strconv.Atoi(r)
	if err != nil {
		return nil
	}
	return l.Limit(limit)
}

func ORDER_ASCEND(l * gorm.DB, r string) *gorm.DB {
	return l.Order(r)
}

func ORDER_DESCEND(l * gorm.DB, r string) *gorm.DB {
	return l.Order(r + " desc")
}

func contains(slice []string, value string) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}
	return false
}