package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/promql/parser"
)

func parse(input string) ([][]*labels.Matcher, error) {
	expr, err := parser.ParseExpr(input)
	if err != nil {
		return nil, err
	}

	return parser.ExtractSelectors(expr), nil
}

func main() {
	var metricNamesOnly = flag.Bool("nameonly", false, "only print metric names")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			continue
		}
		out, err := parse(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "parse error:", err, scanner.Text())
			os.Exit(1)
			return
		}
		if *metricNamesOnly {
			for i := range out {
				for j := range out[i] {
					if out[i][j].Name == `__name__` {
						fmt.Println(out[i][j].Value)
					}
				}
			}
			continue
		}

		for i := range out {
			fmt.Println(out[i])
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
