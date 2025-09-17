import { useState } from "react";
import { markAttendance } from "../../api/teacher";

const Attendance = () => {
  const [form, setForm] = useState({ studentID: "", date: "", status: "Present" });

  const handleSubmit = async () => {
    await markAttendance(form);
    alert("Attendance marked");
  };

  return (
    <div>
      <h2>Mark Attendance</h2>
      <input placeholder="Student ID" onChange={(e) => setForm({ ...form, studentID: e.target.value })} />
      <input type="date" onChange={(e) => setForm({ ...form, date: e.target.value })} />
      <select onChange={(e) => setForm({ ...form, status: e.target.value })}>
        <option value="Present">Present</option>
        <option value="Absent">Absent</option>
      </select>
      <button onClick={handleSubmit}>Submit</button>
    </div>
  );
};

export default Attendance;
