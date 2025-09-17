import { useEffect, useState } from "react";
import { getPerformance } from "../../api/student";

const Performance = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    getPerformance().then((res) => setData(res.data));
  }, []);

  return (
    <div>
      <h2>Performance</h2>
      <ul>
        {data.map((item, i) => (
          <li key={i}>{item.subject}: {item.score}</li>
        ))}
      </ul>
    </div>
  );
};

export default Performance;
