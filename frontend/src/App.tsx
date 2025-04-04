import { todoApi } from "@/api/todo";
import "@/App.css";
import { Todo } from "@/types/todo";
import { useEffect, useState } from "react";

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTodoTitle, setNewTodoTitle] = useState("");
  const [error, setError] = useState<string>("");
  const [isLoading, setIsLoading] = useState(true);
  const [editingTodo, setEditingTodo] = useState<{
    id: number;
    title: string;
  } | null>(null);

  const fetchTodos = async () => {
    try {
      setIsLoading(true);
      const data = await todoApi.list();
      console.log("Fetched todos:", data);
      setTodos(data);
    } catch (error) {
      setError(`Todoの取得に失敗しました: ${(error as Error).message}`);
      console.error("Error fetching todos:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleCreateTodo = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!newTodoTitle.trim()) return;

    try {
      setError("");
      await todoApi.create(newTodoTitle);
      setNewTodoTitle("");
      await fetchTodos();
    } catch (error) {
      setError(`Todoの作成に失敗しました: ${(error as Error).message}`);
      console.error("Error creating todo:", error);
    }
  };

  const handleDeleteTodo = async (id: number) => {
    try {
      setError("");
      await todoApi.delete(id);
      await fetchTodos();
    } catch (error) {
      setError(`Todoの削除に失敗しました: ${(error as Error).message}`);
      console.error("Error deleting todo:", error);
    }
  };

  const handleUpdateTodo = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!editingTodo || !editingTodo.title.trim()) return;

    try {
      setError("");
      await todoApi.update(editingTodo.id, editingTodo.title);
      setEditingTodo(null);
      await fetchTodos();
    } catch (error) {
      setError(`Todoの更新に失敗しました: ${(error as Error).message}`);
      console.error("Error updating todo:", error);
    }
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  return (
    <div className="App">
      <h1>Todoリスト</h1>

      <form onSubmit={handleCreateTodo}>
        <input
          type="text"
          value={newTodoTitle}
          onChange={(e) => setNewTodoTitle(e.target.value)}
          placeholder="新しいTodoを入力"
          required
        />
        <button type="submit" style={{ marginLeft: "10px" }}>
          追加
        </button>
      </form>

      {error && (
        <p
          style={{ color: "red", padding: "10px", backgroundColor: "#ffebee" }}
        >
          {error}
        </p>
      )}

      {isLoading ? (
        <p>読み込み中...</p>
      ) : (
        <ul style={{ listStyle: "none", padding: 0 }}>
          {todos && todos.length > 0 ? (
            todos.map((todo) => (
              <li
                key={todo.id}
                style={{
                  margin: "10px 0",
                  padding: "10px",
                  border: "1px solid #ddd",
                }}
              >
                {editingTodo?.id === todo.id ? (
                  <form
                    onSubmit={handleUpdateTodo}
                    style={{ display: "inline" }}
                  >
                    <input
                      type="text"
                      value={editingTodo.title}
                      onChange={(e) =>
                        setEditingTodo({
                          ...editingTodo,
                          title: e.target.value,
                        })
                      }
                      required
                    />
                    <button type="submit">更新</button>
                    <button type="button" onClick={() => setEditingTodo(null)}>
                      キャンセル
                    </button>
                  </form>
                ) : (
                  <>
                    {todo.title}
                    <button
                      onClick={() =>
                        setEditingTodo({ id: todo.id, title: todo.title })
                      }
                      style={{ marginLeft: "10px" }}
                    >
                      編集
                    </button>
                    <button
                      onClick={() => handleDeleteTodo(todo.id)}
                      style={{ marginLeft: "10px" }}
                    >
                      削除
                    </button>
                  </>
                )}
              </li>
            ))
          ) : (
            <p>Todoがありません</p>
          )}
        </ul>
      )}
    </div>
  );
}

export default App;
