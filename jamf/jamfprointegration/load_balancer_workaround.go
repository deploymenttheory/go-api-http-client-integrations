package jamfprointegration

import (
	"fmt"
	"net/http"
	"strings"
	"time"
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
	j.Sugar.Debugf("Chosen Cookie:%v ", chosenCookie)
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
	j.Sugar.Debug("Starting load balancer workaround")
	var outList []string
	var err error
	var req *http.Request
	var resp *http.Response
	var iterations int

	iterations = 0
	startTimeEpoch := time.Now().Unix()
	endTimeEpoch := startTimeEpoch + int64(LoadBalancerTimeOut.Seconds())

	for i := time.Now().Unix(); i < endTimeEpoch; {
		req, err = http.NewRequest("GET", urlString, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %v", err)
		}

		// Auth required on login screen or 404
		err = j.PrepRequestParamsAndAuth(req)
		if err != nil {
			return nil, fmt.Errorf("error populating auth: %v", err)
		}

		resp, err = j.http.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error sending req: %v", err)
		}

		respCookies := resp.Cookies()

		for _, v := range respCookies {
			if v.Name == LoadBalancerTargetCookie {
				strippedCookie := strings.TrimSpace(v.Value)
				outList = append(outList, strippedCookie)
			}
		}

		uniqueMap := make(map[string]bool)

		for _, str := range outList {
			uniqueMap[str] = true
		}

		cookieDupesRemoved := make([]string, 0, len(uniqueMap))

		for str := range uniqueMap {
			cookieDupesRemoved = append(cookieDupesRemoved, str)
		}

		if len(cookieDupesRemoved) > 1 {
			j.Sugar.Debugf("### COMPLETED LOADBALANCER WORKAROUND ### Dupes removed: %v, outlist: %v", cookieDupesRemoved, outList)
			break
		}

		i = time.Now().Unix()
		iterations += 1
	}
	return &outList, nil

}
