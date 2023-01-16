package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	mongotesting "coolcar/shared/mongo/testing"
	"os"
	"testing"
)
var mongoURI string;
func TestCreatTrip(t *testing.T) {
	mongoURI = "mongodb://localhost:27017"
	c := context.Background()
	mc, err := mongotesting.NewDefaultClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	t.Errorf("111")
	m := NewMongo(mc.Database("coolcar"));
	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: "account1",
		CarId: "112",
		Start: &rentalpb.LocationStatus{
			PoiName: "startPoint",
			Location: &rentalpb.Location{
				Latitude: 30,
				Longitude: 120,
			},
		},
		End:  &rentalpb.LocationStatus{
			PoiName: "endPoint",
			FeeCent: 10000,
			KmDriven: 35,
			Location: &rentalpb.Location{
				Latitude: 35,
				Longitude: 115,
			},
		},
	})
	if err != nil {
		t.Errorf("cannot creat CreateTrip %v", err)
	}
	t.Errorf("1223, %+v", tr)
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}