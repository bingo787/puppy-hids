package main

import "gopkg.in/mgo.v2"

const (
	mongodeploy string = "172.40.219.253:27017"
	db = "agent"
)

var DB *mgo.Database
var err error

