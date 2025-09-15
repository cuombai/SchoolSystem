package admin

type CreateUserRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"` // "student" or "teacher"
    Phone    string `json:"phone,omitempty"`
}

type AuditLog struct {
    Action    string `json:"action"`
    ActorID   string `json:"actor_id"`
    Timestamp string `json:"timestamp"`
    Details   string `json:"details"`
}
