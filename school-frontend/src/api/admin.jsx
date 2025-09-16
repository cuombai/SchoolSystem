import api from "./axios";

export const createUser = (data) => api.post("/admin/user", data);
export const deleteUser = (id) => api.delete(`/admin/user/${id}`);
export const getAuditLogs = () => api.get("/admin/audit");