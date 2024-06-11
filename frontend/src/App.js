import './App.css';
import Login from './components/Login';
import 'bootstrap/dist/css/bootstrap.min.css'; 



function App() {

  // Define your theme object (using Bootstrap CSS variables)
  const yourThemeObject = {
      '--bs-primary': '#black',
      '--bs-secondary': '#white',
      // Add other Bootstrap variable overrides as needed
  };


  return (
      <div style={{
          '--bs-primary': yourThemeObject['--bs-primary'],
          '--bs-secondary': yourThemeObject['--bs-secondary'],
          // Add other Bootstrap variable overrides as needed
      }}>
          <Login />
      </div>
  );
}

export default App;