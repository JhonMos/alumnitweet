package bd

import (
	"context"
	"time"

	"github.com/JhonMos/alumnitweet/models"
)

/*BorroRelacion borra la relación entre usuarios*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("alumnitweet")
	col := db.Collection("relationships")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, err
}
