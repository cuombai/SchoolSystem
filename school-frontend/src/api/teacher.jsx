import api from "./axios";

export const markAttendance = (data) => api.post("/teacher/attendance", data);
export const uploadGrades = (data) => api.post("/teacher/grades", data);
export const uploadRemarks = (data) => api.post("/teacher/remarks", data);
export const getClassList = (classId) => api.get(`/teacher/class/${classId}`);
