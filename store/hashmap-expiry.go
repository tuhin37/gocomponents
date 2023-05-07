package store

/*
	1. impliments hashmap with expiry.
	2. this creates two maps, one for the actual data and one for the expiry epoch
	3. when a key is set, it is set in both the maps
	4. when a key is get, it is first checked in the expiry map, if expiry_epoch is smaller than current epoch, then it is deleted from both the maps and nil is returned
	5. there is a go routines which runs every 1 minute  and deletes all the expired keys from both the maps
*/
