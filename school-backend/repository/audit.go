package repository

import (
	// "context"
	"schoolsystem/school-backend/config"
	"schoolsystem/school-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertAuditLog(log models.AuditLog) error{
	ctx, cancel := config.GetContext()
	defer cancel()

	collection := getDB().Collection("audit_logs")
	_, err := collection.InsertOne(ctx, log)
	return err 
}

func GetAuditLogs(limit int64) ([]models.AuditLog, error){
	ctx, cancel := config.GetContext()
	defer cancel()

	collection := getDB().Collection("audit_logs")

	opts := options.Find().SetSort(bson.D{{"timestamp", -1}}).SetLimit(limit)
	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []models.AuditLog
	if err := cursor.All(ctx, &logs); err != nil {
		return nil, err
	}

	return logs, nil
}