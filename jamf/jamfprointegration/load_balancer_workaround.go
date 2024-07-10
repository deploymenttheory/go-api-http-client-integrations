package jamfprointegration

import (
	"fmt"
	"net/http"
	"slices"
)

// GetSessionCookies retrieves all cookies from the current session
func (j *Integration) getSessionCookies(urlString string) ([]*http.Cookie, error) {
	var returnList []*http.Cookie
	balancerValue, err := j.GetLoadBalancer(urlString)
	if err != nil {
		return nil, fmt.Errorf("error getting load balancer: %v", err)
	}
	returnList = append(returnList, &http.Cookie{Name: LoadBalancerTargetCookie, Value: balancerValue})
	return returnList, nil
}

// GetLoadBalancer programatically always returns the most alphabetical load balancer from a session
func (j *Integration) GetLoadBalancer(urlString string) (string, error) {
	allBalancers, err := j.getAllLoadBalancers(urlString)
	if err != nil {
		return "", fmt.Errorf("error getting load balancers: %v", err)
	}

	chosenCookie := chooseMostAlphabeticalString(*allBalancers)
	return chosenCookie, nil
}

// chooseMostAlphabeticalString returns the most alphabetical string from a list of strings
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

// TODO migrate strings
func (j *Integration) getAllLoadBalancers(urlString string) (*[]string, error) {

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

		for _, v := range respCookies {
			if v.Name == LoadBalancerTargetCookie {
				outList = append(outList, v.Value)
			}
		}

	}

	slices.Sort(outList)
	newList := slices.Compact(outList)
	return &newList, nil

}
