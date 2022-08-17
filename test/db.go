package test

import "gopkg.in/khaiql/dbcleaner.v2"

var defaultCountry = "pk"
var dbCleaners = make(map[string]dbcleaner.DbCleaner)

// DBCleaner is cleaner for default database (pablo_pk)
// Deprecated: all tests MUST be written to work concurrently. You cannot have exclusive access to the DB.
var DBCleaner dbcleaner.DbCleaner
