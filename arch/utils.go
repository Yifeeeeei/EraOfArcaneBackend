package arch

import "github.com/Yifeeeeei/EraOfArcaneBackend/arch/model"

// tools
// see if a string is in a list
func InStringList(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

// see if an id is in a list
func InIdList(id model.Id, list []model.Id) bool {
	for _, i := range list {
		if i.SameAs(id) {
			return true
		}
	}
	return false
}
