import { useEffect, useState } from "react";
import { getTranscript } from "../../api/student";



const Transcript = () => {
  const [transcript, setTranscript] = useState([]);

  useEffect(() => {
    getTranscript().then((res) => setTranscript(res.data));
  }, []);

  return (
    <div>
      <h2>Transcript</h2>
      <ul>
        {transcript.map((item, i) => (
          <li key={i}>{item.term}: {item.grade}</li>
        ))}
      </ul>
    </div>
  );
};

export default Transcript;
