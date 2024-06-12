package main

import (
	"MVC_member/Model"
	"fmt"
)

func main() {
	Model.InsertMember(1, "Iksan", "Surabaya", 90)
	Model.InsertMember(2, "Irfan", "Sidoarjo", 89)
	Model.InsertMember(3, "Ryan", "Sidoarjo", 85)
	Model.InsertMember(4, "Dimas", "Kediri", 95)

	Model.SearchMember(4)

	SearchValue := Model.CariMember(2)
	fmt.Println("Func CariMember :", SearchValue.Data)

	Model.UpdateMember(3, "Gianluki", "Sidoarjo")
	Model.UpdateMember(4, "Renaldy", "Surabaya")
	Model.DeleteMember(2)

	members := Model.ReadAllMember()
	fmt.Println("--------------------------")
	if members != nil {
		for members.Next != nil {
			fmt.Println(members.Next.Data)
			members = members.Next
		}
	}
	fmt.Println("--------------------------")
}
