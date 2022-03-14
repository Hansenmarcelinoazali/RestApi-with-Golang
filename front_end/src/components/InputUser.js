import React, { useEffect, useState } from "react";
import Axios from "axios";

export default function InputUser() {
  const [user, setUser] = useState("");
  const [users, setUsers] = useState([]);
  const apiUrl = "http://localhost:1323/users";

  const submitUser = async () => {
    console.log(user);
    const res = await Axios.post(apiUrl, { username: user });
    console.log(res);
  };

  useEffect(() => {
    const getUsers = async () => {
      const res = await Axios.get(apiUrl);
      if (res != null) {
        setUsers(res.data.data);
      }
    };
    getUsers();
  }, []);

  const deleteUser = async (id) => {
    const res = await Axios.delete(`${apiUrl}/${id}`, )

  }

  if (users == []) {
    return "loading";
  }

  return (
    <div>
      <div className="container-fluid">
        <div className="row">
          {users.map((value) => (
            <div className="col-3">
              <div className="card bg-dark p-3 my-2">
                <h2 className="font-weight-bold bg-dark text-light mx-2">
                  {value.username},{" "}
                </h2>
                <button className="btn bg-light col-3 text-weight-bold">delete</button>
              </div>
            </div>
          ))}
        </div>
      </div>
      <form className="form-group container p-4">
        <label className="font-weight-bold">Input Name</label>
        <input
          className="form-control m-2"
          type="text"
          id="typeURL"
          class="form-control"
          onChange={(e) => setUser(e.target.value)}
        />
        <button className="btn btn-warning my-2" onClick={() => submitUser()}>
          submit
        </button>
      </form>
    </div>
  );
}
