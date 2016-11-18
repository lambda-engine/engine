package analytics

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/satori/go.uuid"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func TrackEvent(r *http.Request, propertyID, category, action, label string, value *int) error {
	if propertyID == "" {
		return errors.New("analytics: GA_TRACKING_ID attribute is missing")
	}
	if category == "" || action == "" {
		return errors.New("analytics: category and action are required")
	}

	v := url.Values{
		"v":   {"1"},
		"tid": {propertyID},
		"ds":  {"app"},
		// "geoid" : {region},
		// "ul" : {language},

		// Anonymously identifies a particular user. See the parameter guide for
		// details:
		// https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#cid
		//
		// Depending on your application, this might want to be associated with the
		// user in a cookie.
		"cid": {uuid.NewV4().String()},
		"t":   {"event"},
		"ec":  {category},
		"ea":  {action},
		"ua":  {r.UserAgent()},
	}

	if label != "" {
		v.Set("el", label)
	}

	if value != nil {
		v.Set("ev", fmt.Sprintf("%d", *value))
	}

	if remoteIP, _, err := net.SplitHostPort(r.RemoteAddr); err != nil {
		v.Set("uip", remoteIP)
	}

	// post the request to GA
	client := urlfetch.Client(appengine.NewContext(r))
	resp, err := client.PostForm("https://www.google-analytics.com/collect", v)
	defer resp.Body.Close()

	// NOTE: Google Analytics returns a 200, even if the request is malformed.
	return err
}
