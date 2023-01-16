package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	mgutil "coolcar/shared/mongo"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// const openIDField = "open_id"

// Mongo defines a mongo dao.
//
type Mongo struct {
	col *mongo.Collection
}

// TripRecord defines a trip record in mongo db.
type TripRecord struct {
	mgutil.IDField        `bson:"inline"`
	mgutil.UpdatedAtField `bson:"inline"`
	Trip                  *rentalpb.Trip `bson:"trip"`
}
// NewMongo creates a mongo dao.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("trip"),
	}
}

// CreateTrip creates a trip.
func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	r := &TripRecord{
		Trip: trip,
	}
	r.ID = mgutil.NewObjID()
	r.UpdatedAt = mgutil.UpdatedAt()

	_, err := m.col.InsertOne(c, r)
	if err != nil {
		log.Fatalf("111 %v", err)
		return nil, err
	}

	return r, nil
}