package sys

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"rmonitor/internals/pkg/models/database"
	"time"
)

type Host struct {
	BaseModel
	HostName string `json:"hostName" bson:"hostName,omitempty"`
	BmcIp    string `json:"bmcIp" bson:"bmcIp,omitempty"`
	BmcMac   string `json:"bmcMac" bson:"bmcMac,omitempty"`
	BmcUser  string `json:"bmcUser" bson:"bmcUser,omitempty"`
	BmcPwd   string `json:"bmcPwd" bson:"bmcPwd,omitempty"`
}

func (r Host) List() (hosts []Host, err error) {
	mdb, err := database.NewMDBDefault()
	if err != nil {
		return nil, err
	}
	table := mdb.DB.Collection("host")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if cur, err := table.Find(ctx, bson.M{}); err != nil {
		return nil, err
	} else {
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var host Host
			err = cur.Decode(&host)
			if err != nil {
				return nil, err
			}
			hosts = append(hosts, host)
		}
	}
	return hosts, err
}