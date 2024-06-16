import './App.css';
import Login from './components/Login';
import 'bootstrap/dist/css/bootstrap.min.css'; 



function App() {

  const yourThemeObject = {
      '--bs-primary': '#black',
      '--bs-secondary': '#white',
  };


  return (
          <Login />
  );
}

export default App;