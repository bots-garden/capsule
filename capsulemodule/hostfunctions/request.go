package hostfunctions

type Request struct {
	Body    string
	Headers map[string]string
	Uri     string
	Method  string
}

/* TODO:
- add url
- GetHeader method
- QueryString
- ...
*/
