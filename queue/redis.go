package queue

/*
	Its a wrapper around redis list.

	Properties:
		1. Name
		2. TTL (in seconds, default: 0, no expiry)

	Methods:
		1. Initialize (name, ttl)
		2. Push (string)
		3. Pop (string)
		4. Count
		5. Delete
		6. DeleteAll
		7. PushAny	=> convert struct to string and store
		8. PopAny	=> convert string to struct and return
		9. PushAnyCompressed => convert struct to string and compress and store
		10. PopAnyCompressed => convert string to struct and decompress and return
		18. export Json
		19. export compressed json
		20. export binary
		18. import Json
		19. import compressed json
		20. import binary

*/
