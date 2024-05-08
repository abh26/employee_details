package crud_functions

import (
	"context"
	"fmt"
	"hardwaremonitioringexporter/pb/pb"
	"log"
	"math"
)

func (d *DatabaseManager) ListEmployeeRecords(ctx context.Context, in *pb.EmployeeRecordsRequest) (*pb.EmployeeRecordsResponse, error) {
	var Count int32
	result := d.Db.Raw(`Select Count(1) from clients`).Find(&Count)
	log.Println("result : ", result)
	var EmployeeDetails []Employee
	Error := d.Db.Raw(`Select * from Employee limit=? and offset=?`, in.Limit, fmt.Sprint((in.Page-1)*in.Limit)).Find(&EmployeeDetails).Error
	if Error != nil {
		log.Println("Error while fetching data : ", Error)
		return &pb.EmployeeRecordsResponse{Status: "failed", Statuscode: "500", Message: "Data could not be fetched"}, nil

	}
	var Data []*pb.Record
	for _, j := range EmployeeDetails {
		var Dummy *pb.Record = &pb.Record{}
		Dummy.Id = int32(j.ID)
		Dummy.Name = j.Name
		Dummy.Position = j.Position
		Dummy.Salary = float32(j.Salary)
		Data = append(Data, Dummy)
	}
	var Page *pb.Page = &pb.Page{
		CurrentPage: in.Page,
		Start:       ((in.Page - 1) * in.Limit) + 1,
		LastPage:    int32(math.Ceil(float64(Count) / float64(in.Limit))),
		PerPage:     in.Limit,
		End:         int32(Count),
		Total:       int32(Count),
	}
	return &pb.EmployeeRecordsResponse{Status: "success", Statuscode: "200", Message: "Successsfully fetched data", Pagination: Page, Data: Data}, nil
}
