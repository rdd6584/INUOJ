package httprsp

func isNotNull(x paramInfo) bool {
	if x.Type == 0 {
		if x.Value == "-1" {
			return true
		}
	} else if x.Type == 1 {
		if x.Value == "" {
			return true
		}
	}
	return false
}
func addParam(str *string, x paramInfo) {
	if x.Type == 0 {
		*str += x.Name + "=" + x.Value + " "
	} else if x.Type == 1 {
		*str += x.Name + "='" + x.Value + "' "
	}
}
func makeWhere(x ...paramInfo) string {
	ret := ""
	fst := true
	for _, cur := range x {
		if isNotNull(cur) {
			if fst {
				ret += "where "
			} else {
				ret += "and "
			}
			addParam(&ret, cur)
		}
	}
	return ret
}
