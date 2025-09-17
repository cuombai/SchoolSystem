import { useEffect, useState } from "react";
import { getAuditLogs } from "../../api/admin";
// import { formatDate } from "@/utils/formatDate";



const AuditLogs = () => {
  const [logs, setLogs] = useState([]);

  useEffect(() => {
    getAuditLogs().then((res) => setLogs(res.data));
  }, []);

  return (
    <div>
      <h2>Audit Logs</h2>
      <ul>
        {logs.map((log, i) => (
          <li>{log.timestamp} - {log.action} by {log.actorID}</li>

        ))}
      </ul>
    </div>
  );
};

export default AuditLogs;
