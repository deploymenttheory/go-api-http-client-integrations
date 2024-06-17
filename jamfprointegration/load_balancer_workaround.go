package jamfprointegration

import (
	"fmt"
	"net/http"
	"slices"
)

func (j *Integration) GetLoadBalancer(urlString string) (string, error) {
	allBalancers, err := j.getAllLoadBalancers(urlString)
	if err != nil {
		return "", fmt.Errorf("error getting load balancers: %v", err)
	}

	chosenCookie := chooseMostAlphabeticalString(*allBalancers)
	return chosenCookie, nil
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

func (j *Integration) getAllLoadBalancers(urlString string) (*[]string, error) {
	var cookiesList []*http.Cookie
	var outList []string
	var err error
	var req *http.Request
	var resp *http.Response
	client := http.Client{}

	for i := 0; i < LoadBalancerPollCount; i++ {
		req, err = http.NewRequest("GET", urlString, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %v", err)
		}

		err = j.PrepRequestParamsAndAuth(req)
		if err != nil {
			return nil, fmt.Errorf("error populating auth: %v", err)
		}

		resp, err = client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error sending req: %v", err)
		}

		respCookies := resp.Cookies()

		for i, v := range respCookies {
			if v.Name == LoadBalancerTargetCookie {
				cookiesList = append(cookiesList, respCookies[i])
				outList = append(outList, v.Value)
			}
		}

	}

	slices.Sort(outList)
	newList := slices.Compact(outList)
	return &newList, nil

}
