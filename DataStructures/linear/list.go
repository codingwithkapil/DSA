package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List

	//insert the element in the front of the list
	Front := intList.PushFront(11) //list [11]

	//inset the element in the back of the list
	Remove := intList.PushBack(23) //list [11, 23]
	intList.PushBack(44)           //list [11, 23, 44]

	intList.InsertAfter(10, Front)    //list [11, 10, 23, 44]
	intList.InsertBefore(12, Front)   //list [12, 11, 10, 23, 44]
	intList.MoveAfter(Front, Remove)  //list [12, 10, 23, 11, 44]
	intList.MoveBefore(Remove, Front) //list [12, 10, 23, 11, 44]
	intList.Remove(Remove)            //list [12, 10, 11, 44]

	//chack if the list's next element is empty is or not
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
