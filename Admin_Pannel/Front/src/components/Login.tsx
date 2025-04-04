import React, { useState } from 'react';

interface LoginProps {
  onLogin: (login: string, password: string) => void;
}

const Login: React.FC<LoginProps> = ({ onLogin }) => {
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onLogin(login, password);
  };

  return (
    <div className="login-container">
      <h2>Вход в систему</h2>
      <form onSubmit={handleSubmit} className="login-form">
        <div className="input-group">
          <input
            type="text"
            value={login}
            onChange={(e) => setLogin(e.target.value)}
            placeholder="Логин"
            required
            autoComplete="username"
          />
        </div>
        <div className="input-group">
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Пароль"
            required
            autoComplete="current-password"
          />
        </div>
        <button type="submit">Войти</button>
      </form>
    </div>
  );
};

export default Login; 