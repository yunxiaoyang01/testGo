package main

import (
	"github.com/globalsign/mgo/bson"
	"regexp"
	"strings"
)

func main()  {
	s:="5ecf7078c4eb5b000bf88108,5ed05ed5db19478e4c6ff0cc,5ed05dd1db19478e4c6fcdda,5ecf72d8c4eb5b000ef8f0b0,5ecbc06dc4eb5b000ff91817"
	idArray :=decodeObjectIDs(s)
	println(idArray)
}
func decodeObjectIDs(s string) []bson.ObjectId {
	var r []bson.ObjectId
	for _, id := range regexp.MustCompile("[,，、]").Split(s, len(s)) {
		objectID := strings.TrimSpace(id)
		if objectID == "" {
			continue
		}
		if !bson.IsObjectIdHex(objectID) {
			println("shibai")
		}
		r = append(r, bson.ObjectIdHex(objectID))
	}
	return r
}
