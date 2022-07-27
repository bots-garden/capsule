package hostfunctions

import "strconv"

// to not display an error code at the end of the error message, code == 0
func CreateStringError(message string, code int) string {
    return "[ERR]["+strconv.Itoa(code)+"]:"+message
    // "[ERR][200]:hello world"
    // message: e.split("]:")[1]
    // code: e.split("]")[1].split("[")[1]
}
