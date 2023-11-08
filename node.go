package handin4mutualexclusion

type Node struct {
	id                int
	port              int
	inCriticalSection bool
}

func EnterRequest (Id int) {
	//wait
	//ask if others are in critsect
	//if one is true, wait and ask again over and over until they're not
	//if all are false, set yourself to true and go in babyh
	//call exit once done
}

func Exit (Id int) {
	//set inCriticalSection to false
	//start over with EnterRequest
}