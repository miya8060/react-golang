import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [message, setMessage] = useState<string>("");
  const [error, setError] = useState<string>("");

  useEffect(() => {
    fetch("http://localhost:8080/api/health")
      .then((response) => {
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
      })
      .then((data) => setMessage(data.message))
      .catch((error) => setError(`エラーが発生しました: ${error.message}`));
  }, []);

  return (
    <div className="App">
      <h1>フロント・バックエンド疎通確認</h1>
      {message && <p>バックエンドからのメッセージ: {message}</p>}
      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  );
}

export default App;
