package hostfunctions

import (
    "github.com/bots-garden/capsule/commons"
    "strings"
)

type Request struct {
    Body    string
    Headers map[string]string
    Uri     string
    Method  string
}

func (req Request) ParseQueryString() map[string]string {
    uriParts := strings.Split(req.Uri, "/?")
    if len(uriParts) > 1 {
        queryString := uriParts[1]
        slice := commons.CreateSliceFromString(queryString, "&")
        if len(slice) > 0 {
            params := commons.CreateMapFromSlice(slice, "=")
            if len(params) > 0 {
                return params
            } else {
                return nil
            }
        } else {
            return nil
        }

    } else {
        return nil
    }
}

/* TODO:
- add url
- GetHeader method
- QueryString
- ...
*/
