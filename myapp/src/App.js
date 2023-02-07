import axios from "axios";
import React, { useState } from "react";
import "./App.css";

function App() {
  const [resdata, setdata] = useState({
    title: "",
    des: "",
  });
  const getPort = () => {
    const url = `${window.location.host}`;
    const pattern = /\d+/g;
    const PORT = url.match(pattern)[0];
    return PORT;
  };
  function action() {
    axios
      .get(`${window.location.href.replace(getPort(), 8000)}`)
      .then(function (response) {
        setdata({
          title: response.data.title,
          des: response.data.description,
        });
      })
      .catch(function (error) {
        // console.log(error);
      });
  }
  return (
    <div
      className="App"
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        position: "relative",
        width: "100%",
        height: "100vh",
        flexDirection: "column",
      }}
    >
      <h3>Title : {resdata.title}</h3>
      <h3>description : {resdata.des}</h3>
      <button
        onClick={() => {
          action();
        }}
        style={{
          position: "relative",
        }}
      >
        action
      </button>
    </div>
  );
}

export default App;
