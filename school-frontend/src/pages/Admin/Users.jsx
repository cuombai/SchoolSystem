import { useState, useEffect } from "react";
import { createUser } from "../../api/admin";
import "../styles/form.css";

const Users = () => {
  const [form, setForm] = useState({ name: "", email: "", role: "teacher", password: "" });

  const handleCreate = async () => {
    try {
      await createUser(form);
      alert("User created");
    } catch (err) {
      alert("Error creating user");
    }
  };

  return (
    <div>
      <h2>Create User</h2>
      <input placeholder="Name" onChange={(e) => setForm({ ...form, name: e.target.value })} />
      <input placeholder="Email" onChange={(e) => setForm({ ...form, email: e.target.value })} />
      <input placeholder="Password" onChange={(e) => setForm({ ...form, password: e.target.value })} />
      <select onChange={(e) => setForm({ ...form, role: e.target.value })}>
        <option value="teacher">Teacher</option>
        <option value="student">Student</option>
      </select>
      <button onClick={handleCreate}>Create</button>
    </div>
  );
};

export default Users;
