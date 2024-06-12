package Model

import (
	"MVC_member/Database"
	"MVC_member/Node"
	"fmt"
)

func InsertMember(id int, nama string, alamat string, point float32) {
	newDataMember := Node.TableMember{
		Data: Node.Member{Id: id, Nama: nama, Alamat: alamat, Point: point},
		Next: nil,
	}
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if Database.DBmember.Next == nil {
		Database.DBmember.Next = &newDataMember
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataMember
	}
}

func ReadAllMember() *Node.TableMember {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if Database.DBmember.Next == nil {
		return nil
	} else {
		return tempLL
	}
}

func SearchMember(id int) {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	cek := false
	if Database.DBmember.Next == nil {
		fmt.Println("DATA MEMBER KOSONG!")
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				fmt.Println("--------------------------")
				fmt.Println("Nomor Member  :", tempLL.Data.Id)
				fmt.Println("Nama Member   :", tempLL.Data.Nama)
				fmt.Println("Alamat Member :", tempLL.Data.Alamat)
				fmt.Println("Point Member  :", tempLL.Data.Point)
				fmt.Println("--------------------------")
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			fmt.Println("ID Tidak Ditemukan!")
		}
	}
}

func CariMember(id int) *Node.TableMember {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	cek := false
	if Database.DBmember.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				return tempLL
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}

func UpdateMember(id int, nama string, alamat string) {
	member := CariMember(id)
	if member != nil {
		member.Data.Nama = nama
		member.Data.Alamat = alamat
		fmt.Println("Update Data Ke", id, "Berhasil")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func DeleteMember(id int) {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if tempLL.Next == nil {
		fmt.Println("Data Kosong")
	} else {
		for tempLL.Next != nil {
			//fmt.Println(tempLL.Next.Data)
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				fmt.Println("Data Ke", id, "berhasil dihapus")
				return
			}
			tempLL = tempLL.Next
		}
	}
}
