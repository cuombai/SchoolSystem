import api from "./axios";

export const getPerformance = () => api.get("/student/performance");
export const getTranscript = () => api.get("/student/transcript");
export const getFeeSummary = () => api.get("/student/fees");
export const getInvoice = (term) => api.get(`/student/invoice/${term}`);
