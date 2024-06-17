package jamfprointegration

import "fmt"

func getAllLoadBalancers(url string, pollCount int) *[]string {
	for i := 0; i < pollCount+1; i++ {
		fmt.Printf("iter-{%v}", i)
	}
	return nil
}

func chooseMostAlphabeticalString(strings []string) string {
	if len(strings) == 0 {
		return ""
	}

	mostAlphabeticalStr := strings[0]
	for _, str := range strings[1:] {
		if str < mostAlphabeticalStr {
			mostAlphabeticalStr = str
		}
	}

	return mostAlphabeticalStr
}
