import { useEffect, useState } from 'react'
import './App.css'

type CalculatedPackage = {
  packSize: number
  count: number
}

function App() {
  const [configuredPackSizes, setConfiguredPackSizes] = useState<number[]>([])
  const [packSizes, setPackSizes] = useState<CalculatedPackage[]>([])
  const [count, setCount] = useState<number | null>(null)
  const [newPackSizes, setNewPackSizes] = useState<string>('')

  useEffect(() => {
    fetchPackSizes()
  }, [])

  const fetchPackSizes = async () => {
    const response = await fetch('http://localhost:8080/packs')
    const data = await response.json()
    setConfiguredPackSizes(data)
  }

  const handleCalculate = async () => {
    const response = await fetch('http://localhost:8080/calculate?items=' + count, {
      method: 'GET',
    })
    const data = await response.json()
    setPackSizes(data)
  }

  const handleUpdatePackSizes = async () => {
    const sizes = newPackSizes.split(',').map(Number)
    await fetch('http://localhost:8080/packs', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ packSizes: sizes }),
    })
    setNewPackSizes('')
    fetchPackSizes()
  }

  return (
    <div className="app">
      <h1>Pack Calculator</h1>
      <div className="card">
        <h2>Configured Pack Sizes</h2>
        <ul style={{ listStyleType: 'none', padding: 0 }}>
          {configuredPackSizes.map(size => (
            <li key={size} style={{paddingBottom: 5, fontSize: 18}}>{size}</li>
          ))}
        </ul>
      </div>
      <div className="card">
        <h2>Calculate Required Packs</h2>
        <input
          type="number"
          onChange={(e) => setCount(Number(e.target.value))}
          placeholder="Enter required items"
        />
        <button onClick={handleCalculate}>Calculate</button>
        {packSizes.map((pack) => (
          <div key={pack.packSize} style={{display: 'flex', justifyContent: 'center', alignItems: 'center'}}>
            <span>{pack.count} x {pack.packSize}</span>
          </div>
        ))}
      </div>
      <div className="card">
        <h2>Update Pack Sizes</h2>
        <input
          type="text"
          value={newPackSizes}
          onChange={(e) => setNewPackSizes(e.target.value)}
          placeholder="Enter pack sizes (comma-separated)"
          style={{width: '50%'}}
        />
        <button onClick={handleUpdatePackSizes}>Update</button>
      </div>
    </div>
  )
}

export default App