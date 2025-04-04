import { useState} from 'react'
import Login from './components/Login'
import DataTable from './components/DataTable'
import './styles.css'

const App = () => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(() => {
    const authData = localStorage.getItem('authData')
    if (!authData) return false
    
    const { expiresAt } = JSON.parse(authData)
    if (new Date().getTime() > expiresAt) {
      localStorage.removeItem('authData')
      return false
    }
    
    return true
  })
  
  const [data, setData] = useState<any[]>([])
  const [isLoading, setIsLoading] = useState(false)

  const handleLogin = (login: string, password: string) => {
    if (login === 'Aidana' && password === 'admin') {
      const expiresAt = new Date().getTime() + (24 * 60 * 60 * 1000) // 24 hours
      localStorage.setItem('authData', JSON.stringify({
        login,
        expiresAt
      }))
      setIsAuthenticated(true)
    } else {
      alert('Неверный логин или пароль')
    }
  }

  const handleLogout = () => {
    localStorage.removeItem('authData')
    setIsAuthenticated(false)
    setData([])
  }

  const fetchData = async () => {
    try {
      setIsLoading(true)
      const response = await fetch('http://176.123.178.135:8080/posts', {
        headers: {
          'Authorization': 'Basic QWlkYW5hOmFkbWlu',
          'Access-Control-Allow-Origin': '*',
        },
      })
      if (!response.ok) {
        throw new Error('Ошибка загрузки данных')
      }
      const result = await response.json()
      setData(result)
    } catch (error) {
      alert('Ошибка загрузки данных')
      console.error(error)
    } finally {
      setIsLoading(false)
    }
  }

  if (!isAuthenticated) {
    return <Login onLogin={handleLogin} />
  }

  return (
    <div className="container">
      <div className="header">
        <h1>MET GALA</h1>
        <button className="logout-button" onClick={handleLogout}>Выйти</button>
      </div>
      <div className="buttons">
        <button onClick={fetchData}>Показать данные</button>
      </div>
      {isLoading ? (
        <div className="loading">Загрузка данных...</div>
      ) : (
        <>
          {data.length > 0 ? (
            <DataTable data={data} />
          ) : (
            <div className="no-data">
              Нажмите кнопку "Показать данные" для загрузки
            </div>
          )}
        </>
      )}
    </div>
  )
}

export default App
