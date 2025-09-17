import { useState } from "react";
import { uploadGrades } from "../../api/teacher";

const Grades = () => {
  const [form, setForm] = useState({ studentID: "", subject: "", score: "" });

  const handleSubmit = async () => {
    await uploadGrades(form);
    alert("Grades uploaded");
  };

  return (
    <div>
      <h2>Upload Grades</h2>
      <input placeholder="Student ID" onChange={(e) => setForm({ ...form, studentID: e.target.value })} />
      <input placeholder="Subject" onChange={(e) => setForm({ ...form, subject: e.target.value })} />
      <input placeholder="Score" onChange={(e) => setForm({ ...form, score: e.target.value })} />
      <button onClick={handleSubmit}>Submit</button>
    </div>
  );
};

export default Grades;
