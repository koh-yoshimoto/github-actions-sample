import logo from "./github-actions-logo.png";
import "./App.css";
import React, { useState } from "react";
import axios from "axios";

function App() {
  axios.defaults.headers.post["Content-Type"] =
    "application/json;charset=utf-8";

  const replyPlaceholder = "reply from lambda is displayed here";

  const [name, setName] = useState("");
  const [reply, setReply] = useState(replyPlaceholder);

  const url =
    "https://zy7cemqd63cv3pzn4rlqbnk2ge0pgglk.lambda-url.ap-northeast-1.on.aws/";

  const onButtonClicked = () => {
    const data = { name: name };
    axios
      .post(url, data)
      .then(({ data }) => {
        setReply(data.message);
        setName("");
      })
      .catch(() => {
        console.log("failed");
        setReply(replyPlaceholder);
        setName("");
      });
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          <code>input your name</code>
        </p>
        <input
          className="App-form"
          type="text"
          onChange={(event) => setName(event.target.value)}
          value={name}
        ></input>
        <button onClick={onButtonClicked} className="App-button">
          <code>submit</code>
        </button>
        <p>
          <code>{reply}</code>
        </p>

        <a
          className="App-link"
          href="https://docs.github.com/ja/actions"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn CI/CD by GitHub Actions!
        </a>
      </header>
    </div>
  );
}

export default App;
