import { useNavigate } from "react-router-dom";

function useLogout(setToken) { 
  const navigate = useNavigate();

  const handleLogout = () => {
    setToken(null);

    navigate("/"); 
  };

  return { handleLogout };
}

export default useLogout;