package admin

import (
    // "schoolsystem/school-backend/models"
    "schoolsystem/school-backend/repository"
)

func GetAuditLogs() ([]AuditLog, error) {
    logs, err := repository.GetAuditLogs()
    if err != nil {
        return nil, err
    }

    var result []AuditLog
    for _, l := range logs {
        result = append(result, AuditLog{
            Action:    l.Action,
            ActorID:   l.ActorID,
            Timestamp: l.Timestamp.Format("2006-01-02 15:04"),
            Details:   l.Details,
        })
    }
    return result, nil
}
