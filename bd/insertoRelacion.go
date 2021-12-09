package bd

import (
	"context"
	"time"

	"github.com/JhonMos/alumnitweet/models"
)

/*InsertoRelacion graba la relación en la BD*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("alumnitweet")
	col := db.Collection("relationships")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, err
}
