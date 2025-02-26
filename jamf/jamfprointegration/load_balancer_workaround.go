package jamfprointegration

import (
	"fmt"
	"net/http"
	"slices"
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
	j.Sugar.Debug("LOGHERE")
	j.Sugar.Debug("Starting cookie magic")
	var outList []string
	var err error
	var req *http.Request
	var resp *http.Response

	startTimeEpoch := time.Now().Unix()
	j.Sugar.Debugf("Start time: %d", startTimeEpoch)
	endTimeEpoch := startTimeEpoch + int64(LoadBalancerTimeOut.Seconds())
	j.Sugar.Debugf("End Time: %d", endTimeEpoch)

	for i := time.Now().Unix(); i < endTimeEpoch; i++ {
		req, err = http.NewRequest("GET", urlString, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %v", err)
		}

		// Auth required on login screen or 404
		err = j.PrepRequestParamsAndAuth(req)
		if err != nil {
			return nil, fmt.Errorf("error populating auth: %v", err)
		}

		resp, err = j.httpExecutor.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error sending req: %v", err)
		}

		respCookies := resp.Cookies()
		j.Sugar.Debugf("Cookies got: %+v", respCookies)

		for _, v := range respCookies {
			if v.Name == LoadBalancerTargetCookie {
				j.Sugar.Debugf("Appending: %v", v.Value)
				outList = append(outList, v.Value)
			}
		}
		j.Sugar.Debugf("BEGIN DUPE REMOVAL. OUTLIST: %v", outList)
		cookieDupesRemoved := slices.Compact(outList)
		j.Sugar.Debugf("DUPES REMOVED: %v", cookieDupesRemoved)
		if len(cookieDupesRemoved) > 1 {
			break
		}

	}
	return &outList, nil

}
