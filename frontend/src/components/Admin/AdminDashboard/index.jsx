import React, { useState, useEffect } from 'react';
import './index.css';
import { useAuth } from '../../../hooks/AuthContext';
import { useTranslation } from 'react-i18next';
import { useAxiosWithToken } from '../../../utils/api';
import "react-toastify/dist/ReactToastify.css";
import CustomAdminNavbar from '../AdminNavbar/index';

function AdminDashboard({ setToken }) {

    const { t } = useTranslation();
    const [error, setError] = useState(null);
    const { token } = useAuth();
    const api = useAxiosWithToken();
    useEffect(() => {
        
        const fetchData = async () => {
            if(token !== null){
            try {
                const response = await api.get("admin");
                console.log(response);
                if (response.status === 200) {
                    console.log("has authorization");
                } else {
                    setError(response);
                }
            } catch (error) {
                setError(error);
            }
        }
        };
        fetchData();
    }, [token]);
    
    return ( 
        <>
        <CustomAdminNavbar setToken={setToken} t={t}></CustomAdminNavbar>
        <div></div>
        </>
    )
}

export default AdminDashboard;