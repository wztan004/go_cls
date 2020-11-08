// Package constants contain the project-level configuration
package constants

const (
	Timeout = "30m" // max session duration before it times out
	Address = ":5221" // port and IP address

	LogFile = "confidential/log.txt" // location of logs
	UserFile = `confidential/users.csv` // location of user data
	LatestMthLess1 = `confidential/venues_202009.csv` // location of venue data for previous month
	LatestMth = `confidential/venues_202010.csv` // location of venue data for current month
	Cert = `confidential/cert.pem` // location of SSH certificate
	Key = `confidential/key.pem` // location of SSH key
)
