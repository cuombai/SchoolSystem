import { useEffect, useState } from "react";
import { getFeeSummary } from "../../api/student";

const Fees = () => {
  const [fees, setFees] = useState([]);

  useEffect(() => {
    getFeeSummary().then((res) => setFees(res.data));
  }, []);

  return (
    <div>
      <h2>Fee Summary</h2>
      <ul>
        {fees.map((item, i) => (
          <li key={i}>{item.term}: {item.amount} KES</li>
        ))}
      </ul>
    </div>
  );
};

export default Fees;
