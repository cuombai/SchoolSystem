import { useEffect, useState } from "react";
import { getClassList } from "../../api/teacher";

const ClassList = () => {
  const [students, setStudents] = useState([]);

  useEffect(() => {
    getClassList("classID").then((res) => setStudents(res.data));
  }, []);

  return (
    <div>
      <h2>Class List</h2>
      <ul>
        {students.map((s) => (
          <li key={s._id}>{s.name}</li>
        ))}
      </ul>
    </div>
  );
};

export default ClassList;
